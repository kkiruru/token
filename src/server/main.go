package main

import(
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func main(){
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":1324"))
}

func login(c echo.Context) error {
	id := c.FormValue("id")
	password := c.FormValue("password")

	if confirmLoginInfo(id, password){
		accessToken, refreshToken := createToken(id)
		tokens := make(map[string]string)

		tokens["accessToken"] = accessToken
		tokens["refreshToken"] = refreshToken

		return c.String(http.StatusOK, makeResponse(tokens))
	}else{
		return c.String(http.StatusUnauthorized, "invalid id or password!!")
	}
}


/*
	JSON Type의 response를 리턴한다
	key-value에 추가적인 정보를 더 한다
*/
func makeResponse(values map[string]string) string{
	jsonString, _ := json.Marshal(values)
	return string(jsonString)
}




/*
	TODO
	id와 password로 사용자를 인증한다.
	그냥 test 비교만 수행*/
func confirmLoginInfo(id string, password string) bool {
	return true
}


/*
	TODO
	토큰을 생성한다

	access token에는 사용자 아이디와 토큰 만료 시간이 들어간다
	refresh token에는 토큰 만료 시간이 들어간다
*/
func createToken(id string) (string, string){
	return "alsdkjf19p0aklsjdfasdf1232512", "asdfasdpa;sdfj;laskdjf;"
}