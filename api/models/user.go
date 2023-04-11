package models

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPrimaryKey struct {
	UserId string `json:"user_id"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
