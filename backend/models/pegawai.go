package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// Struktur model Pegawai
type Pegawai struct {
	ID           int            `json:"id"`
	Nama         string         `json:"nama"`
	NIP          string         `json:"nip"`
	TanggalLahir string         `json:"tanggal_lahir"`
	Role         string         `json:"role"`
	Foto         sql.NullString `json:"foto"`
	CreatedAt    time.Time      `json:"created_at"`
}

// Buat tabel pegawai (dengan kolom created_at)
func CreateTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS pegawai (
		id SERIAL PRIMARY KEY,
		nama TEXT NOT NULL,
		nip TEXT UNIQUE NOT NULL,
		tanggal_lahir TEXT NOT NULL,
		role TEXT DEFAULT 'pegawai',
		foto TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("📦 Tabel pegawai siap digunakan")
}

// Seed data awal pegawai (jika belum ada data)
func SeedPegawai(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pegawai").Scan(&count)
	if err != nil {
		fmt.Println("❌ Gagal menghitung pegawai:", err)
		return
	}
	if count > 0 {
		fmt.Println("👥 Pegawai data already seeded")
		return
	}

	names := []string{"Andi", "Budi", "Citra", "Dewi", "Eka", "Fajar", "Gita", "Hadi", "Intan", "Joko"}
	for i, name := range names {
		nip := fmt.Sprintf("19890%04d", rand.Intn(10000))
		tgl := time.Now().AddDate(-20-i, 0, 0).Format("2006-01-02") // Format ISO untuk PostgreSQL
		_, err := db.Exec(`
			INSERT INTO pegawai (nama, nip, tanggal_lahir, role) 
			VALUES ($1, $2, $3, $4)
		`, name, nip, tgl, "pegawai")

		if err != nil {
			fmt.Println("❌ Seed error:", err)
		}
	}
	fmt.Println("✅ Seeded 10 pegawai")
}
