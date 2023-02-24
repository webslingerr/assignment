package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UpdateUserBalance struct {
	Id string `json:"id"`
	Balance int `json:"balance"`
}