package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/repositories"

	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterProductRouters = func(route fiber.Router, db *gorm.DB) {
	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)
	route.Get("", productController.GetAll)
	route.Get("/:id", productController.GetById)
	route.Patch("/:id", productController.Update)
	route.Delete("/:id", productController.Delete)
	route.Post("", productController.Create)
}
