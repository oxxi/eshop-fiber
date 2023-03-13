package entity

import "gorm.io/gorm"

type UserEntity struct {
	*gorm.Model
	Name        string `gorm:"varchar(255);not null"`
	Email       string `gorm:"varchar(255)"`
	PhoneNumber string `gorm:"varchar(255)"`
	Password    []byte `gorm:"not null"`
}

func (UserEntity) TableName() string {
	return "Users"
}
