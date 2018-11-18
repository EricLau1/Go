package models

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {

	/* necessário fazer o instalar o REDIS no SO ~ Leia o arquivo REDIS.md */
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

}
