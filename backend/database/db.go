package database

import (
	"database/sql"
	"fmt"
	"log"
	"monlap/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Load .env jika ada, gunakan env global jika tidak
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found. Using environment variables.")
	}

	// Ambil variabel environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validasi semua variabel penting
	missingVars := []string{}
	if host == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if port == "" {
		missingVars = append(missingVars, "DB_PORT")
	}
	if user == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if password == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
	}
	if dbname == "" {
		missingVars = append(missingVars, "DB_NAME")
	}
	if len(missingVars) > 0 {
		log.Fatalf("❌ Missing required environment variables: %v", missingVars)
	}

	// Buat connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// Inisialisasi koneksi DB
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Cek apakah database bisa dijangkau
	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Database not reachable: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	// Buat tabel secara otomatis jika belum ada
	models.CreateTable(DB)

	// Seed data pegawai jika tabel kosong
	models.SeedPegawai(DB)
}
