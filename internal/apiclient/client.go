package apiclient

//go:generate mockgen -source=./client.go -destination=./mock/client_mock.go -package=mock

import (
	"github.com/liderman/traderstack/internal/domain"
	"github.com/liderman/traderstack/internal/grpc/gen/tinkoff/investapi"
	"github.com/shopspring/decimal"
	"time"
)

type ApiClient interface {
	GetCandles(figi string, from time.Time, to time.Time, interval domain.CandleInterval) ([]*domain.HistoricCandle, error)
	GetShares() ([]*domain.Share, error)
	PostOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error)
	GetLastPrices() ([]*domain.LastPrice, error)
	GetAccounts() ([]*domain.Account, error)
	GetSandboxAccounts() ([]*domain.Account, error)
	OpenSandboxAccount() (string, error)
	GetSandboxPositions(accountId string) (*domain.PositionsResponse, error)
	SandboxPayIn(accountId string, amount *domain.MoneyValue) error
	PostSandboxOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error)
	GetSandboxPortfolio(accountId string) ([]*domain.PortfolioPosition, error)
	GetPortfolio(accountId string) ([]*domain.PortfolioPosition, error)
	GetOrders(accountId string) ([]*domain.OrderState, error)
	GetSandboxOrders(accountId string) ([]*domain.OrderState, error)
}

type RealClient struct {
	cMarket     investapi.MarketDataServiceClient
	cInstrument investapi.InstrumentsServiceClient
	cOrder      investapi.OrdersServiceClient
	cUsers      investapi.UsersServiceClient
	cSandbox    investapi.SandboxServiceClient
	cOperations investapi.OperationsServiceClient
}

func NewRealClient(
	cMarket investapi.MarketDataServiceClient,
	cInstrument investapi.InstrumentsServiceClient,
	cOrder investapi.OrdersServiceClient,
	cUsers investapi.UsersServiceClient,
	cSandbox investapi.SandboxServiceClient,
	cOperations investapi.OperationsServiceClient,
) *RealClient {
	return &RealClient{
		cMarket:     cMarket,
		cInstrument: cInstrument,
		cOrder:      cOrder,
		cUsers:      cUsers,
		cSandbox:    cSandbox,
		cOperations: cOperations,
	}
}
