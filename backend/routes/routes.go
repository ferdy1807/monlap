package routes

import (
	"monlap/controllers"
	"monlap/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("✅ Monlap Backend is running")
	})

	// Login endpoint
	app.Post("/login", controllers.Login)

	// Group route /api dengan proteksi JWT
	api := app.Group("/api", middleware.JWTMiddleware)

	// Endpoint dashboard (terproteksi)
	api.Get("/dashboard", controllers.GetDashboard)

	// Endpoint ringkasan pegawai (total dan waktu input terakhir)
	api.Get("/pegawai/summary", controllers.GetPegawaiSummary)

	// Endpoint semua data pegawai (jika perlu ditampilkan semuanya)
	api.Get("/pegawai", controllers.GetAllPegawai)
}
