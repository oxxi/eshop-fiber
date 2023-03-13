package models

type UserModel struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}
