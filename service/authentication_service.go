package service

import (
	"errors"

	"github.com/kkiruru/token/model/dao"
	"github.com/kkiruru/token/model/dto"
)

//id password가 맞는지 확인한다
func Authentication(id string, password string) (*dto.User, error) {

	account, err := dao.FindById("id1")

	if err != nil {
		return nil, err
	}

	if account.ID != id || account.Password != password {
		return nil, errors.New("invalid id or password")
	}

	return account, nil
}
