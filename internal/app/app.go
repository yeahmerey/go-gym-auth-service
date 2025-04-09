package app

import (
	"fmt"
	"go-gym-auth-service/internal/app/auth"
	"net/http"
	"go-gym-auth-service/internal/app/config"	
)

func Run(configFile string) error {
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		return fmt.Errorf("ошибка при загрузке конфигурации: %v", err)
	}

	fmt.Printf("Запуск HTTP сервера на порту: %s\n", cfg.HTTPServer.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth service is running"))
	})

	mux.HandleFunc("/register", auth.RegisterHandler)

	mux.HandleFunc("/login", auth.LoginHandler)

	err = http.ListenAndServe(":"+cfg.HTTPServer.Port, mux)
	if err != nil {
		return fmt.Errorf("ошибка при запуске HTTP сервера: %v", err)
	}

	return nil
}
