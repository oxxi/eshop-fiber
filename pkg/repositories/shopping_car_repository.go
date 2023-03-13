package repositories

import (
	"github.com/oxxi/eshop/pkg/entity"
	"gorm.io/gorm"
)

type IShoppingCarRepository interface {
	AddToCar(entity *entity.ShoppingCarEntity) error
	DeleteFromCar(entity *entity.ShoppingCarEntity) error
}

type shoppingCarRepository struct {
	Db *gorm.DB
}

// AddToCar implements IShoppingCarRepository
func (r *shoppingCarRepository) AddToCar(entity *entity.ShoppingCarEntity) error {
	panic("unimplemented")
}

// DeleteFromCar implements IShoppingCarRepository
func (r *shoppingCarRepository) DeleteFromCar(entity *entity.ShoppingCarEntity) error {
	panic("unimplemented")
}

func NewShoppingCarRepository(db *gorm.DB) IShoppingCarRepository {
	return &shoppingCarRepository{
		Db: db,
	}
}
