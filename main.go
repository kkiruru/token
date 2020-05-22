package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kkiruru/token/model/dto"
	"github.com/kkiruru/token/service"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login)
	log.Fatal(router.Run(":8080"))
}

func login(c *gin.Context) {
	var login dto.User
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//for TEST Only...
	account, err := service.Authentication(login.ID, login.Password)
	if err == nil {
		token, err := service.CreateToken(account.ID)
		if err == nil {
			c.JSON(http.StatusOK, token)
			return
		}
	}

	c.JSON(http.StatusUnprocessableEntity, err.Error())
}
