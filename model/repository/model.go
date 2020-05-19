package repository

type DataBase struct {
	data map[string]string
}

var repo map[string][]string

func New() {
	repo := new(DataBase)
	repo = create(repo)
}

func create(repo *DataBase) *DataBase {
	repo.data["user1"] = "pass1"
	repo.data["user2"] = "pass2"
	repo.data["user3"] = "pass3"
	repo.data["user4"] = "pass4"

	return repo
}
