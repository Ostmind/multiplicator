package config

import (
	"errors"
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type ServerConfig struct {
	Host                  string        `env:"SERVER_HOST"             envDefault:"localhost"`
	Port                  int           `env:"SERVER_PORT"             envDefault:"64333"`
	ServerReadTimeout     time.Duration `env:"SERVER_READ_TIMEOUT"     envDefault:"5s"`
	ServerWriteTimeout    time.Duration `env:"SERVER_WRITE_TIMEOUT"    envDefault:"5s"`
	ServerIdleTimeout     time.Duration `env:"SERVER_IDLE_TIMEOUT"     envDefault:"10s"`
	ServerShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10s"`
	EnvType               string        `env:"ENV_TYPE"                envDefault:"local"`
}

func MustNew() *ServerConfig {
	cfgEnv := ServerConfig{}

	err := env.Parse(&cfgEnv)
	if err != nil {
		log.Fatalf("err loading file: %s", err)
	}

	errs := cfgEnv.Validate()
	if errs != nil {
		log.Fatalf("err validating config: %s", errs.Error())
	}

	return &cfgEnv
}

func (cfg *ServerConfig) Validate() (result error) {
	if cfg.Host == "" {
		result = errors.Join(result, ErrNoServerHost)
	}

	if cfg.Port == 0 {
		result = errors.Join(result, ErrNoServerPort)
	}

	return result
}
