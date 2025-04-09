package app

import (
	"context"
	"auth-service/internal/app/config"
	"auth-service/internal/app/connections"
	"auth-service/internal/app/start"
	"auth-service/pkg/logger"
)

func Run(configFiles ...string) {
	ctx := context.Background()

	cfg, err := config.New(configFiles...)
	if err != nil {
		panic(err)
	}

	logger.Init()

	conn, err := connections.New(cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start.HTTP(ctx, cfg, conn)
}
