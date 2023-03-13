package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/repositories"
	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterShoppingCarRouter = func(route fiber.Router, db *gorm.DB) {

	shoppingCarRepository := repositories.NewShoppingCarRepository(db)
	shoppingCarService := service.NewShoppingCarService(shoppingCarRepository)
	shoppingController := controllers.NewShoppingCarController(shoppingCarService)

	route.Get("", shoppingController.GetAll)
	route.Post("", shoppingController.Add)
	route.Delete("", shoppingController.Delete)
}
