package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/middleware"
	"github.com/oxxi/eshop/pkg/repositories"
	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterShoppingCarRouter = func(route fiber.Router, db *gorm.DB) {

	shoppingCarRepository := repositories.NewShoppingCarRepository(db)
	shoppingCarService := service.NewShoppingCarService(shoppingCarRepository)
	shoppingController := controllers.NewShoppingCarController(shoppingCarService)

	route.Get("", middleware.DeserializeUser, shoppingController.GetAll)
	route.Post("", middleware.DeserializeUser, shoppingController.Add)
	route.Delete("", middleware.DeserializeUser, shoppingController.Delete)
}
