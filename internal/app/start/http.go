package start

import (
	"context"
	"fmt"
	"net/http"

	"auth-service/internal/app/config"
	"auth-service/internal/app/connections"
)

func HTTP(ctx context.Context, cfg *config.Config, conn *connections.Connections) {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	fmt.Println("Starting server at", cfg.HTTPServer.Address)
	err := http.ListenAndServe(cfg.HTTPServer.Address, mux)
	if err != nil {
		panic(err)
	}
}
