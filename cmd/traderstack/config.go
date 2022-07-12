package main

import "time"

// Config Application configuration
type Config struct {
	LogLevel              string        `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"TS_LOG_LEVEL" default:"info"`
	LogJSON               bool          `long:"log-json" description:"Enable force log format JSON" env:"TS_LOG_JSON"`
	ListenHTTP            string        `long:"listen-http" description:"Listening host:port for HTTP-server" env:"TS_LISTEN_HTTP" default:":8080"`
	BrokerAPIToken        string        `long:"broker-api-token" description:"Production token for access broker API" env:"TS_BROKER_API_TOKEN"`
	BrokerAPITimeout      time.Duration `long:"broker-api-timeout" description:"Broker API timeout" env:"TS_BROKER_API_TIMEOUT" default:"5s"`
	BrokerSandboxAPIToken string        `long:"broker-sandbox-api-token" description:"Sandbox token for access broker API" env:"TS_BROKER_SANDBOX_API_TOKEN"`
	ProductionMode        bool          `long:"production-mode" description:"Use production API token" env:"TS_PRODUCTION_MODE"`
	DataDir               string        `long:"data-dir" description:"Directory path to save application data" env:"TS_DATA_DIR" default:"data"`
	UseDevProxy           bool          `long:"use-dev-proxy" description:"Enable proxy for static files" env:"TS_USE_DEV_PROXY"`
	DevProxyHost          string        `long:"dev-proxy-host" description:"host:port for static files dev server" env:"TS_DEV_PROXY_HOST" default:"127.0.0.1:3000"`
	StaticFilesDir        string        `long:"static-files-dir" description:"Path for static files" env:"TS_STATIC_FILES_DIR" default:"../../web/dist"`
}
