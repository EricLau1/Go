package models

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound  = errors.New("usuário não encontrado.")
	ErrInvalidLogin  = errors.New("Login inválido.")
	ErrUsernameTaken = errors.New("Username ja existe.")
)

type User struct {
	id int64
}

func NewUser(username string, hash []byte) (*User, error) {

	// verifica se o username existe
	exists, err := client.HExists("user:by-username", username).Result()

	if exists {
		return nil, ErrUsernameTaken
	}

	id, err := client.Incr("user:next-id").Result()

	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user:%d", id)

	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "username", username)
	pipe.HSet(key, "password", hash)
	pipe.HSet("user:by-username", username, id)

	_, err = pipe.Exec()

	if err != nil {

		return nil, err

	}

	return &User{id}, nil
}

func (user *User) GetUserId() (int64, error) {

	return user.id, nil

}

func (user *User) GetUsername() (string, error) {

	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(key, "username").Result()

}

func (user *User) GetHashPassword() ([]byte, error) {

	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(key, "password").Bytes()
}

func (user *User) Authenticate(password string) error {

	hash, err := user.GetHashPassword()

	if err != nil {

		return err

	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {

		return ErrInvalidLogin

	}

	return err

}

func GetUserById(id int64) (*User, error) {

	return &User{id}, nil

}

func GetUserByUsername(username string) (*User, error) {

	id, err := client.HGet("user:by-username", username).Int64()

	if err == redis.Nil {

		return nil, ErrUserNotFound

	} else if err != nil {

		return nil, err

	}

	return GetUserById(id)
}

func AuthenticateUser(username, password string) (*User, error) {

	user, err := GetUserByUsername(username)

	if err != nil {

		return nil, err

	}

	return user, user.Authenticate(password)
}

func RegisterUser(username, password string) error {

	cost := bcrypt.DefaultCost

	// gera um hash para a senha
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return err
	}

	_, err = NewUser(username, hash)

	return err

}
