package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterLoginRoutes = func(route fiber.Router, db *gorm.DB) {
	loginService := service.NewLoginService()
	loginControler := controllers.NewLoginController(loginService)

	route.Post("", loginControler.LogIn)
	route.Get("/logout", loginControler.LogOut)
}
