package middleware

import (
	"errors"
	"fmt"
	"monlap/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token tidak ditemukan di header Authorization",
		})
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Format token harus: Bearer <token>",
		})
	}

	tokenString := parts[1]

	// Parse dan validasi token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Validasi algoritma
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("algoritma token tidak didukung: %v", t.Header["alg"])
		}
		return utils.SecretKey, nil
	})

	// Tangani error token
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token telah kedaluwarsa",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token tidak valid",
		})
	}

	// Validasi token sah
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token tidak sah",
		})
	}

	// Simpan token ke context untuk digunakan di handler
	c.Locals("user", token)
	return c.Next()
}
