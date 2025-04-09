package main

import (
	"flag"
	"go-gym-auth-service/internal/app"
	"log"
)

func main() {

	configFile := flag.String("config", "./internal/app/.env", "Path to config file")
	flag.Parse()

	err := app.Run(*configFile)
	if err != nil {
		log.Fatalf("Ошибка при запуске приложения: %v", err)
	}
}
