package controllers

import (
	"database/sql"
	"log"
	"monlap/database"
	"monlap/models"

	"github.com/gofiber/fiber/v2"
)

// GET /pegawai - Mengambil semua data pegawai
func GetAllPegawai(c *fiber.Ctx) error {
	rows, err := database.DB.Query(`
		SELECT id, nama, nip, tanggal_lahir, role, foto, created_at 
		FROM pegawai
	`)
	if err != nil {
		log.Println("❌ Query pegawai gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data dari database",
		})
	}
	defer rows.Close()

	var list []fiber.Map

	for rows.Next() {
		var p models.Pegawai

		err := rows.Scan(
			&p.ID,
			&p.Nama,
			&p.NIP,
			&p.TanggalLahir,
			&p.Role,
			&p.Foto,
			&p.CreatedAt,
		)
		if err != nil {
			log.Println("❌ Scan data pegawai gagal:", err)
			continue
		}

		list = append(list, fiber.Map{
			"id":            p.ID,
			"nama":          p.Nama,
			"nip":           p.NIP,
			"tanggal_lahir": p.TanggalLahir,
			"role":          p.Role,
			"foto":          nullStringToString(p.Foto),
			"created_at":    p.CreatedAt,
		})
	}

	return c.JSON(list)
}

// GET /pegawai/summary - Mengambil ringkasan data pegawai
func GetPegawaiSummary(c *fiber.Ctx) error {
	var total int
	var lastInput sql.NullString

	err := database.DB.QueryRow(`
		SELECT COUNT(*), COALESCE(MAX(created_at::TEXT), '') FROM pegawai
	`).Scan(&total, &lastInput)
	if err != nil {
		log.Println("❌ Query summary pegawai gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data pegawai",
		})
	}

	return c.JSON(fiber.Map{
		"total_pegawai":     total,
		"terakhir_ditambah": lastInput.String,
	})
}

// Helper untuk menghindari null pointer saat foto NULL
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
