package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/service"
)

type IShoppingCarController interface {
	GetAll(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
}

type shoppingCarController struct {
	shoppingService service.IShoppingCarService
}

// Add implements IShoppingCarController
func (s *shoppingCarController) Add(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// Delete implements IShoppingCarController
func (s *shoppingCarController) Delete(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// GetAll implements IShoppingCarController
func (s *shoppingCarController) GetAll(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func NewShoppingCarController(s service.IShoppingCarService) IShoppingCarController {
	return &shoppingCarController{
		shoppingService: s,
	}
}
