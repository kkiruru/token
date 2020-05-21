package dao

import (
	//"github.com/kkiruru/token/model/dto"
	"github.com/kkiruru/token/model/repository"
)

func UserDao(id string) {
	//var user = dto.User{}

	_, _ = repository.LoadTable("account", "id1")
}
