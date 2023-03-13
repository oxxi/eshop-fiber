package entity

type ShoppingCarEntity struct {
	ProductId      int64
	Quantity       int64
	ProductoEntity ProductoEntity `gorm:"foreignKey:ProductId"`
}

func (ShoppingCarEntity) TableName() string {
	return "ShoppingCars"
}
