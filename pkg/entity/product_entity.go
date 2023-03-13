package entity

import (
	"github.com/oxxi/eshop/pkg/models"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type ProductoEntity struct {
	*gorm.Model
	Id       uint   `gorm:"primarykey"`
	Name     string `gorm:"varchar(255)"`
	Detail   string `gorm:"varchar(255)"`
	Price    float64
	Quantity int64
}

func (ProductoEntity) TableName() string {
	return "Productos"
}

func (e ProductoEntity) FromEntity() models.ProductModel {
	return models.ProductModel{
		Id:          e.Id,
		Name:        e.Name,
		Description: e.Detail,
		Quantity:    uint64(e.Quantity),
		Price:       e.Price,
	}
}

func ToEntity(m models.ProductModel) ProductoEntity {
	return ProductoEntity{
		Id:       m.Id,
		Name:     m.Name,
		Detail:   m.Description,
		Price:    m.Price,
		Quantity: int64(m.Quantity),
	}
}
