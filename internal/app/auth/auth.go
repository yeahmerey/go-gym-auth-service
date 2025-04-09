package auth

import (
	"errors"
	"sync"
)

type User struct {
	Username string
	Password string
}

var userStore = struct {
	sync.RWMutex
	users map[string]User
}{users: make(map[string]User)}

func AddUser(username, password string) error {
	userStore.Lock()
	defer userStore.Unlock()

	if _, exists := userStore.users[username]; exists {
		return errors.New("пользователь уже существует")
	}

	userStore.users[username] = User{Username: username, Password: password}
	return nil
}

func ValidateUser(username, password string) bool {
	userStore.RLock()
	defer userStore.RUnlock()

	user, exists := userStore.users[username]
	if !exists {
		return false
	}

	return user.Password == password
}
