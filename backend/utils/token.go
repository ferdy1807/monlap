package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey untuk JWT, diambil dari .env
var SecretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken membuat JWT dengan menyimpan ID dan NIP pegawai
func GenerateToken(id int, nip string) (string, error) {
	// Durasi token berlaku: 24 jam
	expirationTime := time.Now().Add(24 * time.Hour)

	// Payload token
	claims := jwt.MapClaims{
		"id":  id,
		"nip": nip,
		"exp": expirationTime.Unix(),
	}

	// Buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dengan secret key
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
