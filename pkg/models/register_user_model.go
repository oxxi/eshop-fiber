package models

type RegisterUserModel struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type LoginUserModel struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
