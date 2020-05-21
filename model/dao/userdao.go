package dao

import (
	"encoding/json"
	"errors"
	//"github.com/kkiruru/token/model/dto"
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

	return nil, errors.New("not found user info")
}

func FindAll() (*[]dto.User, error) {
	response, err := repository.LoadTable("account")

	if err != nil {
		return nil, err
	}

	var slice []dto.User

	for _, item := range response.([]interface{}) {

		user := dto.User{}

		bodyBytes, _ := json.Marshal(item)
		json.Unmarshal(bodyBytes, &user)
		//
		//
		//item.
		//user := item.(dto.User)
		slice = append(slice, user)
	}
	//
	//for _, item := range response {
	//	fmt.Printf("%v", item.(map[string]interface{})["title"])
	//}
	//
	//data := i.([]dto.User)

	return &slice, nil
}
