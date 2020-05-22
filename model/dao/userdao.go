package dao

import (
	"encoding/json"
	"errors"

	"github.com/kkiruru/token/model/dto"
	"github.com/kkiruru/token/model/repository"
)

func FindById(id string) (*dto.User, error) {
	users, err := FindAll()

	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("Account does not exist")
}

func FindAll() (*[]dto.User, error) {
	table, err := repository.LoadTable("account")

	if err != nil {
		return nil, err
	}

	var users []dto.User

	for _, record := range table.([]interface{}) {

		user := dto.User{}

		bodyBytes, _ := json.Marshal(record)
		json.Unmarshal(bodyBytes, &user)
		users = append(users, user)
	}

	return &users, nil
}
