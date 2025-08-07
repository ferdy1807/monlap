package utils

import (
	"log"
	"os"
)

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  []byte
	AppEnv     string
)

// LoadEnv mengambil semua env var dan validasi
func LoadEnv() {
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	AppEnv = os.Getenv("APP_ENV")
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))

	if DBHost == "" || DBPort == "" || DBUser == "" || DBPassword == "" || DBName == "" || len(JWTSecret) == 0 {
		log.Fatal("❌ Missing required environment variables for database or JWT. Check your .env file.")
	}
}

// IsDev digunakan untuk cek environment development
func IsDev() bool {
	return AppEnv == "DEV"
}
