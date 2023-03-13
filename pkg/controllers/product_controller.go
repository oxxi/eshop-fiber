package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/models"
	"github.com/oxxi/eshop/pkg/service"
)

type IProductController interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type productController struct {
	productService service.IProductService
}

// Create implements IProductController
func (p *productController) Create(c *fiber.Ctx) error {

	productModel := new(models.ProductModel)

	if err := c.BodyParser(productModel); err != nil {
		return err
	}

	model, err := p.productService.Save(*productModel)
	if err != nil {
		return err
	}

	return c.JSON(model)

}

// Delete implements IProductController
func (p *productController) Delete(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := p.productService.Delete(uint64(id)); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

// GetAll implements IProductController
func (p *productController) GetAll(c *fiber.Ctx) error {

	products, _ := p.productService.GetAll()

	c.Response().SetStatusCode(fiber.StatusOK)
	return c.JSON(products)
}

// GetById implements IProductController
func (p *productController) GetById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, _ := p.productService.GetBy(uint64(id))

	c.Response().SetStatusCode(fiber.StatusOK)
	return c.JSON(product)
}

// Update implements IProductController
func (p *productController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	modelProduct := models.ProductModel{}
	c.BodyParser(&modelProduct)
	updatedModel, err := p.productService.Update(uint64(id), modelProduct)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Response().SetStatusCode(fiber.StatusOK)
	return c.JSON(updatedModel)
}

func NewProductController(p service.IProductService) IProductController {
	return &productController{
		productService: p,
	}
}
