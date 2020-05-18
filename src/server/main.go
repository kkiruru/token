package main

import(
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
		return c.String(http.StatusOK, token(id))
	}else{
		return c.String(http.StatusOK, "invalid id or password!!")
	}
}

func confirmLoginInfo(id string, password string) bool {


	return true
}

func token(id string) string{
	return "alsdkjf19p0aklsjdfasdf1232512"
}