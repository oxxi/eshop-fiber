package models

type RegisterUserModel struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type LoginUserModel struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
