package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/kkiruru/token/model/dao"
	"github.com/kkiruru/token/model/dto"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login)
	log.Fatal(router.Run(":8080"))
}

func login(c *gin.Context) {
	var u dto.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//for TEST Only...
	user, _ := dao.FindById("id1")

	if user.ID != u.ID || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userId string) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "golang-token-example")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
