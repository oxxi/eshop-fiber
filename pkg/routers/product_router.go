package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/middleware"
	"github.com/oxxi/eshop/pkg/repositories"

	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterProductRouters = func(route fiber.Router, db *gorm.DB) {
	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)
	route.Get("", middleware.DeserializeUser, productController.GetAll)
	route.Get("/:id", middleware.DeserializeUser, productController.GetById)
	route.Patch("/:id", middleware.DeserializeUser, productController.Update)
	route.Delete("/:id", middleware.DeserializeUser, productController.Delete)
	route.Post("", middleware.DeserializeUser, productController.Create)
}
