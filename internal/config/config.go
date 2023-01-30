package config

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// GRPCAddress is the address to bind the gRPC server to.
	GRPCAddress string

	// HTTPAddress is the address to bind the HTTP server to.
	HTTPAddress string

	// DevMode to indicate development mode. When true, the program would spin up utilities for debugging and
	// provide a more contextual message when encountered a panic. See internal/server/httpserver/http.go for the
	// actual implementation details.
	DevMode bool `split_words:"true"`

	// LimiterEnabled to indicate whether the rate limiter should be enabled.
	LimiterEnabled bool `split_words:"true"`

	// LogLevel is the log level to use. Valid values are "debug", "info", "warn", "error", "fatal", "panic".
	LogLevel string `split_words:"true" default:"info" required:"true"`

	// LogJsonStdout is whether to log JSON logs (instead of pretty-print logs) to stdout for the ease of log collection.
	LogJsonStdout bool `split_words:"true" default:"false"`

	// TracingEnabled to indicate whether to enable OpenTelemetry tracing.
	TracingEnabled bool `split_words:"true"`

	// TracingExporters to indicate which exporters to use for tracing.
	// Valid values are: jaeger, otlp, stdout (for debug).
	TracingExporters []string `split_words:"true" default:"jaeger"`

	// TracingSampleRate to indicate the sampling rate for tracing.
	// Valid values are: 0.0 (disabled), 1.0 (all traces), or a value between 0.0 and 1.0 (sampling rate).
	TracingSampleRate float64 `split_words:"true" default:"1.0"`

	// infrastructure components connection instructions

	// NatsURL is the URL of the NATS server. See https://pkg.go.dev/github.com/nats-io/nats.go#Connect
	// for more information on how to construct a NATS URL.
	NatsURL string `required:"true" split_words:"true" default:"nats://127.0.0.1:4222"`

	// RedisURL is the URL of the Redis server, and by default uses redis db 2, to avoid potential collision
	// with the already running backend instance. See https://pkg.go.dev/github.com/redis/go-redis/v9#ParseURL
	// for more information on how to construct a Redis URL.
	RedisURL string `required:"true" split_words:"true" default:"redis://127.0.0.1:6379/1"`

	// SentryDSN is the DSN of the Sentry server. See https://pkg.go.dev/github.com/getsentry/sentry-go#ClientOptions
	SentryDSN string `split_words:"true"`

	// GRPCServerShutdownTimeout is the timeout for the gRPC server to shut down gracefully.
	GRPCServerShutdownTimeout time.Duration `required:"true" split_words:"true" default:"60s"`
}

func Parse() (*Config, error) {
	var config Config
	err := envconfig.Process("penguin_livehouse", &config)
	if err != nil {
		_ = envconfig.Usage("penguin_livehouse", &config)
		return nil, fmt.Errorf("failed to parse configuration: %w. More info on how to configure this service is located at https://pkg.go.dev/exusiai.dev/livehouse/internal/config#Config", err)
	}

	return &config, nil
}
