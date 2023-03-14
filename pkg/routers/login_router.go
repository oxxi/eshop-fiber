package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/controllers"
	"github.com/oxxi/eshop/pkg/middleware"
	"github.com/oxxi/eshop/pkg/repositories"
	"github.com/oxxi/eshop/pkg/service"
	"gorm.io/gorm"
)

var RegisterLoginRoutes = func(route fiber.Router, db *gorm.DB) {
	loginRepository := repositories.NewUserRepository(db)
	loginService := service.NewLoginService(loginRepository)
	loginControler := controllers.NewLoginController(loginService)

	route.Post("", loginControler.LogIn)
	route.Get("/logout", middleware.DeserializeUser, loginControler.LogOut)
}
