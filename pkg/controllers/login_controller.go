package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oxxi/eshop/pkg/models"
	"github.com/oxxi/eshop/pkg/service"
)

type ILoginController interface {
	LogIn(c *fiber.Ctx) error
	LogOut(c *fiber.Ctx) error
}

type loginController struct {
	loginService service.ILoginService
}

// LogIn implements ILoginController
func (l *loginController) LogIn(c *fiber.Ctx) error {
	//	secret := "secreto"

	user := new(models.LoginUserModel)

	if err := c.BodyParser(user); err != nil {
		return err
	}
	if !l.loginService.LogIn(user) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "secureCookie_"
	cookie.Value = "estoyAutorizado"
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)
	return c.SendStatus(fiber.StatusOK)
}

// LogOut implements ILoginController
func (l *loginController) LogOut(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewLoginController(s service.ILoginService) ILoginController {
	return &loginController{
		loginService: s,
	}
}
