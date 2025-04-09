package auth

import (
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}

	if err := AddUser(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь зарегистрирован"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}

	if !ValidateUser(user.Username, user.Password) {
		http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Успешный вход"))
}
