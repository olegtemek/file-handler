package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Address     string        `env:"HTTP_ADDRESS"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT"`
	Env         string        `env:"ENV"`
	DatabaseUrl string        `env:"DATABASE_URL"`
}

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)

	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}
