package main

import (
	"fmt"
	"go-gym-auth-service/internal/app/auth"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth service is running"))
	})

	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/login", auth.LoginHandler)

	http.Handle("/profile", auth.TokenAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile page"))
	})))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
