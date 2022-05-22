package main

import (
	"context"
	"crypto/tls"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jessevdk/go-flags"
	"github.com/liderman/traderstack/internal/apiclient"
	"github.com/liderman/traderstack/internal/datamanager"
	"github.com/liderman/traderstack/internal/engine"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/liderman/traderstack/internal/grpcsrv"
	infov1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/info/v1"
	stackv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/stack/v1"
	strategyv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/strategy/v1"
	"github.com/liderman/traderstack/internal/stackfuncs"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"
)

var cfg Config
var logger *zap.Logger

func init() {
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	checkError(err)

	var lvl zapcore.Level
	err = lvl.UnmarshalText([]byte(cfg.LogLevel))
	checkError(err)

	opts := zap.NewProductionConfig()
	opts.Level = zap.NewAtomicLevelAt(lvl)
	opts.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if !cfg.LogJSON {
		opts.Encoding = "console"
		opts.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err = opts.Build()
	checkError(err)
}

func main() {
	logger.Info("Starting app")

	if cfg.BrokerSandboxAPIToken == "" {
		logger.Fatal("Укажите токен песочницы (Sandbox) задав переменную окружения TS_BROKER_SANDBOX_API_TOKEN")
	}

	if cfg.BrokerAPIToken == "" {
		logger.Fatal("Укажите production токен задав переменную окружения TS_BROKER_API_TOKEN")
	}

	baseClientOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: 6 * time.Minute, Timeout: 1 * time.Second}),
		grpc.WithBlock(),
	}
	baseClientOpts = appendMetadataGrpcOption(baseClientOpts, "x-app-name", "traderstack")

	sbClientOpts := appendMetadataGrpcOption(baseClientOpts, "Authorization", "Bearer "+cfg.BrokerSandboxAPIToken)
	prodClientOpts := appendMetadataGrpcOption(baseClientOpts, "Authorization", "Bearer "+cfg.BrokerAPIToken)

	investCtx, investCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer investCancel()

	sbInvestApiConn, err := grpc.DialContext(
		investCtx,
		"invest-public-api.tinkoff.ru:443",
		sbClientOpts...,
	)
	if err != nil {
		logger.Fatal("can't connect to SANDBOX invest api", zap.Error(err))
	}

	prodInvestApiConn, err := grpc.DialContext(
		investCtx,
		"invest-public-api.tinkoff.ru:443",
		prodClientOpts...,
	)
	if err != nil {
		logger.Fatal("can't connect to PRODUCTION invest api", zap.Error(err))
	}

	marketApi := investapi.NewMarketDataServiceClient(prodInvestApiConn)
	insApi := investapi.NewInstrumentsServiceClient(prodInvestApiConn)
	orderApi := investapi.NewOrdersServiceClient(prodInvestApiConn)
	usersApi := investapi.NewUsersServiceClient(prodInvestApiConn)
	sandboxApi := investapi.NewSandboxServiceClient(sbInvestApiConn)
	operationsApi := investapi.NewOperationsServiceClient(prodInvestApiConn)
	realApi := apiclient.NewRealClient(marketApi, insApi, orderApi, usersApi, sandboxApi, operationsApi)
	mainCache := cache.New(time.Minute, time.Minute*15)
	insManager := datamanager.NewInstruments(realApi, mainCache)
	marketManager := datamanager.NewMarketData(realApi, mainCache)

	sfr := engine.NewStackFuncRepository()
	sfr.Register(stackfuncs.NewActionBuyMarket(realApi, marketManager, insManager))
	sfr.Register(stackfuncs.NewActionSellMarket(realApi, marketManager, insManager))
	sfr.Register(stackfuncs.NewActionTakeProfit(realApi, marketManager, insManager))
	sfr.Register(stackfuncs.NewActionStopLoss(realApi, marketManager, insManager))
	sfr.Register(stackfuncs.NewPortfolioLots(realApi))
	sfr.Register(stackfuncs.NewInOrdersBuyMarketLots(realApi))
	sfr.Register(stackfuncs.NewInOrdersSellMarketLots(realApi))
	sfr.Register(stackfuncs.NewBoolean())
	sfr.Register(stackfuncs.NewDecimal())
	sfr.Register(stackfuncs.NewFigiByTicker())
	sfr.Register(stackfuncs.NewGe())
	sfr.Register(stackfuncs.NewInteger())
	sfr.Register(stackfuncs.NewLe())
	sfr.Register(stackfuncs.NewRsi(marketManager))
	sfr.Register(stackfuncs.NewString())
	sm := engine.NewStackManager(sfr)

	// Engine
	sEngine := engine.NewStrategyEngine(sm, sfr)
	go func() {
		logger.Info("Running StrategyEngine")
		sEngine.Run()
	}()

	stackSrv := grpcsrv.NewStackServer(sm, sfr, engine.NewTestStack(sm, sfr))
	infoSrv := grpcsrv.NewInfoServer(insManager, realApi)
	engineSrv := grpcsrv.NewStrategyServer(sEngine)
	logger.Info("Starting gRPC server", zap.String("listen", cfg.ListenGRPC))
	srvListen, err := net.Listen("tcp", cfg.ListenGRPC)
	if err != nil {
		logger.Fatal("Failed to listen GRPC server", zap.Error(err))
		return
	}

	recoverFromPanicHandler := func(p interface{}) error {
		err := fmt.Errorf("recovered from panic: %s", p)
		logger.Error("recovered from panic", zap.Error(err))

		return err
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoverFromPanicHandler),
	}

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*15), // 15 MiB
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)

	stackv1.RegisterStackAPIServer(s, stackSrv)
	infov1.RegisterInfoAPIServer(s, infoSrv)
	strategyv1.RegisterStrategyAPIServer(s, engineSrv)
	reflection.Register(s)

	err = s.Serve(srvListen)
	if err != nil {
		logger.Fatal("Stopped gRPC server with error", zap.Error(err))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error init: %s.\nFor help use -h\n", err)
		os.Exit(1)
	}
}
