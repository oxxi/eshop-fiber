package config

import (
	"github.com/oxxi/eshop/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:123456789@tcp(localhost:3306)/Books?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	d.AutoMigrate(&entity.ProductoEntity{}, &entity.ShoppingCarEntity{}, &entity.UserEntity{})

	db = d
}

func GetDB() *gorm.DB {
	return db
}
