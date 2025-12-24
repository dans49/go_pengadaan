package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("RAHASIA_SUPER")

func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	// simpan ke context
	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])
	c.Locals("role", claims["role"])
	c.Locals("nama", claims["nama_lengkap"]) // nama adalah seperti objek, nama_lengkap adalah field diambil dari jwt.go

	return c.Next()
}
