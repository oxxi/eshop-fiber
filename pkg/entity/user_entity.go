package entity

import (
	"fmt"

	"github.com/oxxi/eshop/pkg/models"
	"gorm.io/gorm"
)

type UserEntity struct {
	*gorm.Model
	Name        string `gorm:"varchar(255);not null"`
	UserName    string `gorm:"varchar(255);not null"`
	Email       string `gorm:"varchar(255)"`
	PhoneNumber string `gorm:"varchar(255)"`
	Password    []byte `gorm:"not null"`
}

func (UserEntity) TableName() string {
	return "Users"
}

func (e UserEntity) FromEntity() models.UserModel {
	return models.UserModel{
		Id:       e.ID,
		Name:     e.Name,
		UserName: e.UserName,
	}
}

func ToRegisterUserEnity(m models.RegisterUserModel) UserEntity {
	return UserEntity{
		Model:       &gorm.Model{},
		Name:        fmt.Sprintf("%s %s", m.Name, m.LastName),
		UserName:    m.UserName,
		Email:       m.Email,
		PhoneNumber: m.Phone,
		Password:    []byte(m.Password),
	}
}
