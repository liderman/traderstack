package main

import "time"

// Config Application configuration
type Config struct {
	LogLevel              string        `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"TS_LOG_LEVEL" default:"info"`
	LogJSON               bool          `long:"log-json" description:"Enable force log format JSON" env:"TS_LOG_JSON"`
	ListenGRPC            string        `long:"listen-grpc" description:"Listening host:port for grpc-server" env:"TS_LISTEN_GRPC" default:":9999"`
	BrokerAPIToken        string        `long:"broker-api-token" description:"Production token for access broker API" env:"TS_BROKER_API_TOKEN"`
	BrokerAPITimeout      time.Duration `long:"broker-api-timeout" description:"Broker API timeout" env:"TS_BROKER_API_TIMEOUT" default:"5s"`
	BrokerSandboxAPIToken string        `long:"broker-sandbox-api-token" description:"Sandbox token for access broker API" env:"TS_BROKER_SANDBOX_API_TOKEN"`
	ProductionMode        bool          `long:"production-mode" description:"Use production API token" env:"TS_PRODUCTION_MODE"`
	DataDir               string        `long:"data-dir" description:"Directory path to save application data" env:"TS_DATA_DIR" default:"data"`
}
