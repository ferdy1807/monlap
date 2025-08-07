package main

import (
	"fmt"
	"log"
	"os"

	"monlap/database"
	"monlap/routes"
	"monlap/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file jika ada
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found. Using environment variables.")
	}

	// Validasi dan ambil semua env var
	utils.LoadEnv()

	// Koneksi ke database
	database.Connect()

	// Inisialisasi Fiber
	app := fiber.New()

	// Folder static untuk akses foto
	app.Static("/uploads", "./uploads")

	// Inisialisasi routes
	routes.SetupRoutes(app)

	// Ambil port dari .env atau default 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Monlap backend running on port %s\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
