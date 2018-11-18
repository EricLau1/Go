package models

import (
	"fmt"
	"strconv"
)

type Update struct {
	id int64
}

func NewUpdate(userId int64, body string) (*Update, error) {

	// incrementa o valor existente na chave
	id, err := client.Incr("update:next-id").Result()

	if err != nil {

		return nil, err

	}

	// concatana o id na string e retorna como uma chave do REDIS
	key := fmt.Sprintf("update:%d", id)

	// construindo um HashMap no REDIS
	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "user_id", userId)
	pipe.HSet(key, "body", body)
	pipe.LPush("updates", id)
	pipe.LPush(fmt.Sprintf("user:%d:updates", userId), id)

	_, err = pipe.Exec()

	if err != nil {

		return nil, err

	}

	return &Update{id}, nil
}

func (update *Update) GetBody() (string, error) {

	// Sprintf concatena valores com string
	key := fmt.Sprintf("update:%d", update.id)

	// retorna o resultado de acordo com a chave de busca
	return client.HGet(key, "body").Result()

}

func (update *Update) GetUser() (*User, error) {

	key := fmt.Sprintf("update:%d", update.id)

	userId, err := client.HGet(key, "user_id").Int64()

	if err != nil {

		return nil, err

	}

	return GetUserById(userId)
}

func queryUpdates(key string) ([]*Update, error) {

	// retorna uma lista de até 10 strings
	updateIds, err := client.LRange(key, 0, 10).Result()

	if err != nil {

		return nil, err

	}

	// constroi um vetor de Update de acordo com a quantidade de Ids de updates encontrados no banco
	updates := make([]*Update, len(updateIds))

	for i, strId := range updateIds {

		// converte uma string para inteiro ou retona um erro
		id, err := strconv.Atoi(strId)

		if err != nil {
			return nil, err
		}

		updates[i] = &Update{int64(id)}

	}

	return updates, nil

}

func GetAllUpdates() ([]*Update, error) {

	return queryUpdates("updates")

}

func GetUpdates(userId int64) ([]*Update, error) {

	key := fmt.Sprintf("user:%d:updates", userId)

	return queryUpdates(key)

}

func PostUpdate(userId int64, body string) error {

	// adiciona o comentário ao REDIS
	_, err := NewUpdate(userId, body)

	return err

}
