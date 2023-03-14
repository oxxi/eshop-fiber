package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	errors := models.ValidateStruct(user)
	if errors != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  errors,
		})
	}

	if !l.loginService.LogIn(user) {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "user o password invalid",
		})
	}

	tokenBytes := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenBytes.Claims.(jwt.MapClaims)

	claims["sub"] = 1
	claims["exp"] = now.Add(24 * 1240).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, _ := tokenBytes.SignedString([]byte("my_secret_key"))

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
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
