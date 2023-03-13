package entity

import (
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

/* func (e UserEntity) ToEntity(m models.UserModel) UserEntity {
	return UserEntity{
		Name: m.Name,

	}
}
*/
