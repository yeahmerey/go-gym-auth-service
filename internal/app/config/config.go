package config


import (
	"github.com/caarlos0/env/v6" 
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	HTTPServer HTTPServerConfig
}

type HTTPServerConfig struct {
	Port string `env:"HTTP_PORT" default:"8080"`
}

func NewConfig(filenames ...string) (*Config, error) {
	for _, f := range filenames {
		_ = godotenv.Load(f)
	}

	cfg := &Config{}
	defaults.SetDefaults(cfg)

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
