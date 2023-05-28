package def

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Config struct {
	HTTPPort       string `env:"HTTP_PORT" envDefault:"8080"`
	BasePath       string `env:"HTTP_BASE_PATH" envDefault:""`
	AllowedOrigins string `env:"HTTP_ALLOWED_ORIGINS" envDefault:"*"`
	Host           string `env:"HTTP_HOST"  envDefault:""`
	LogLevel       string `env:"LOG_LEVEL" envDefault:""`
}

func NewConfig(configFiles ...string) (*Config, error) {
	var c Config
	err := godotenv.Load(configFiles...)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return &c, env.Parse(&c, env.Options{
		RequiredIfNoDef: true,
	})

}

const (
	levelInfo    = "info"
	levelWarning = "warning"
	levelError   = "error"
)

func NewLogger(level string) (l *zap.SugaredLogger, err error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true

	switch strings.ToLower(level) {
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case levelInfo:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case levelWarning:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case levelError:
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	basicLogger, err := cfg.Build()
	l = basicLogger.Sugar()

	return
}