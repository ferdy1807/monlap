package controllers

import (
	"database/sql"
	"fmt"
	"monlap/database"
	"monlap/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		NIP          string `json:"nip"`
		TanggalLahir string `json:"tanggal_lahir"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		if utils.IsDev() {
			fmt.Println("❌ Gagal parsing request body:", err)
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format permintaan tidak valid.",
		})
	}

	if req.NIP == "" || req.TanggalLahir == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "NIP dan Tanggal Lahir wajib diisi.",
		})
	}

	// Ambil data dari database berdasarkan NIP
	var (
		id       int
		dbTglLhr string
	)
	err := database.DB.QueryRow(`
		SELECT id, tanggal_lahir 
		FROM pegawai 
		WHERE nip = $1
	`, req.NIP).Scan(&id, &dbTglLhr)

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "NIP atau Tanggal Lahir salah.",
		})
	} else if err != nil {
		if utils.IsDev() {
			fmt.Println("❌ Database error:", err)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan server.",
		})
	}

	if dbTglLhr != req.TanggalLahir {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "NIP atau Tanggal Lahir salah.",
		})
	}

	// Buat token JWT
	token, err := utils.GenerateToken(id, req.NIP)
	if err != nil {
		if utils.IsDev() {
			fmt.Println("❌ Gagal generate token:", err)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan saat membuat token.",
		})
	}

	if utils.IsDev() {
		fmt.Println("✅ Login berhasil untuk ID:", id, "NIP:", req.NIP)
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil.",
		"token":   token,
		// Jika tidak diperlukan di frontend, hapus baris berikut:
		"id":  id,
		"nip": req.NIP,
	})
}
