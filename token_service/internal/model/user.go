package model

type UserDTO struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
