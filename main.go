package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kkiruru/token/model/dto"
	"github.com/kkiruru/token/service"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login)
	router.GET("/api1", api1)
	router.GET("/api2", api2)
	router.GET("/api3", api3)

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
		access_token, _, err := service.CreateToken(account.ID)
		if err == nil {
			c.JSON(http.StatusOK, access_token)
			return
		}
	}

	c.JSON(http.StatusUnprocessableEntity, err.Error())
}

func api1(c *gin.Context) {
	c.JSON(http.StatusOK, "api1")
}

func api2(c *gin.Context) {
	c.JSON(http.StatusOK, "api2")
}

func api3(c *gin.Context) {
	fmt.Println("called api3 >>")
	time.Sleep(time.Second * 3)
	c.JSON(http.StatusOK, "api3")
}
