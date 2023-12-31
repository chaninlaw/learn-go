package fiber

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func checkMiddleWare(c *fiber.Ctx) error {
	url := c.OriginalURL()
	method := c.Method()
	start := time.Now().Format("2006-01-02 15:04:05")

	// JWT ware here
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// fmt.Println("claims", claims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	fmt.Printf("URL = %s, Method = %s,Time = %s\n", url, method, start)

	return c.Next()
}

func superadminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "superadmin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
