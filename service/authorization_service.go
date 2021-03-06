package service

//token을 발급한다

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "golang-token-example")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", "", err
	}

	return token, "refresh_token", nil
}

func TokenAuthorize(token string) error {

	return nil
}

func TokenRefresh(accessToken string, refreshToken string) (string, string, error) {

	err := TokenAuthorize(accessToken)
	if err != nil {
		return "", "", err
	}

	return "", "", nil
}
