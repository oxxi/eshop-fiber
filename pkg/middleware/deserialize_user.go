package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/oxxi/eshop/config"
	"github.com/oxxi/eshop/pkg/repositories"
)

func DeserializeUser(c *fiber.Ctx) error {
	var token string
	authorization := c.Get("Authorization")
	//search token in auth Header or in cookies HttpOnly
	if strings.HasPrefix(authorization, "Bearer ") {
		token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		token = c.Cookies("token")
	}
	//if token no exit return status unauthorized
	if token == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unathirized, user login /v1/api/login",
		})
	}

	//parse JWT
	tokenBytes, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token",
		})
	}
	claims, ok := tokenBytes.Claims.(jwt.MapClaims)
	if !ok || !tokenBytes.Valid {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "invalid token claim"})
	}

	db := config.GetDB()
	userRepo := repositories.NewUserRepository(db)

	id, _ := strconv.ParseInt(fmt.Sprint(claims["sub"]), 0, 0)
	currentUser, err := userRepo.GetById(uint64(id))

	if err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "error while parsing Token"})
	}

	c.Locals("currentUser", currentUser)

	return c.Next()
}
