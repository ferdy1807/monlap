package controllers

import (
	"fmt"
	"monlap/database"
	"monlap/models"
	"monlap/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetDashboard(c *fiber.Ctx) error {
	// Ambil token dari context
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || !userToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token tidak valid",
		})
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Gagal membaca klaim token",
		})
	}

	// Debug hanya saat development
	if utils.IsDev() {
		fmt.Println("🔐 JWT Claims:", claims)
	}

	// Ambil ID dari claims
	idFloat, ok := claims["id"].(float64)
	if !ok {
		if utils.IsDev() {
			fmt.Println("❌ Gagal parsing ID dari token")
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token tidak valid (ID tidak ditemukan)",
		})
	}
	userID := int(idFloat)

	if utils.IsDev() {
		fmt.Println("✅ ID dari token:", userID)
	}

	// Query ke database berdasarkan ID
	var p models.Pegawai
	err := database.DB.QueryRow(`
		SELECT id, nama, nip, tanggal_lahir, role, foto, created_at
		FROM pegawai 
		WHERE id = $1`, userID).
		Scan(&p.ID, &p.Nama, &p.NIP, &p.TanggalLahir, &p.Role, &p.Foto, &p.CreatedAt)

	if err != nil {
		if utils.IsDev() {
			fmt.Println("❌ Query error:", err)
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Pegawai tidak ditemukan",
		})
	}

	if utils.IsDev() {
		fmt.Println("📦 Data pegawai ditemukan:", p)
	}

	// Return data dashboard
	return c.JSON(p)
}
