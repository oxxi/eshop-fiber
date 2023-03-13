package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/config"
	"github.com/oxxi/eshop/pkg/routers"
	"gorm.io/gorm"
)

func setUpRouters(app *fiber.App, Db *gorm.DB) {
	api := app.Group("/api/v1")
	routers.RegisterProductRouters(api.Group("/product"), Db)
	routers.RegisterLoginRoutes(api.Group("/login"), Db)
	routers.RegisterShoppingCarRouter(api.Group("/car"), Db)
}

func main() {

	app := fiber.New()
	config.Connect()
	setUpRouters(app, config.GetDB())

	log.Fatal(app.Listen(":8080"))
}
