package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengubah plaintext password menjadi hashed password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // DefaultCost = 10
	if err != nil {
		if IsDev() {
			fmt.Println("❌ Gagal melakukan hash password:", err)
		}
		return "", err
	}
	return string(hash), nil
}

// CheckPassword memverifikasi password dengan hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if IsDev() {
			fmt.Println("❌ Password tidak cocok:", err)
		}
		return false
	}
	return true
}
