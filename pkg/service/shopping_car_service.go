package service

import (
	"github.com/oxxi/eshop/pkg/models"
	"github.com/oxxi/eshop/pkg/repositories"
)

type IShoppingCarService interface {
	AddItem(item *models.ShoppingCarModel)
	DeleteItem(item *models.ShoppingCarModel)
	CheckOut(car *models.CheckOut)
}

type shoppingCarService struct {
	repository repositories.IShoppingCarRepository
}

// AddItem implements IShoppingCarService
func (s *shoppingCarService) AddItem(item *models.ShoppingCarModel) {
	panic("unimplemented")
}

// CheckOut implements IShoppingCarService
func (s *shoppingCarService) CheckOut(car *models.CheckOut) {
	panic("unimplemented")
}

// DeleteItem implements IShoppingCarService
func (s *shoppingCarService) DeleteItem(item *models.ShoppingCarModel) {
	panic("unimplemented")
}

var shoppingCarInstance *shoppingCarService

func NewShoppingCarService(r repositories.IShoppingCarRepository) IShoppingCarService {
	once.Do(func() {
		shoppingCarInstance = &shoppingCarService{
			repository: r,
		}
	})
	return shoppingCarInstance
}
