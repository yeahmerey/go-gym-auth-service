package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	DB         DBConfig
}

type HTTPServerConfig struct {
	Address string `env:"HTTP_ADDRESS" envDefault:":8080"`
}

type DBConfig struct {
	DSN string `env:"DB_DSN,required"`
}

func New(filenames ...string) (*Config, error) {
	_ = godotenv.Load(filenames...)
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
