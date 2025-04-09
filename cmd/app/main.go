package main

import (
	"flag"
	"go-gym-auth-service/internal/app"
)

func main() {
	configPath := flag.String("config", ".env", "path to config file")
	flag.Parse()

	app.Run(*configPath)

}
