package dto

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type MyJsonName struct {
	Account []struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Username string `json:"username"`
	} `json:"account"`
}
