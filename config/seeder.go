package config

import (
	"Medqueue-Alta-BE/features/user/data"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedAdmin(db *gorm.DB) {
    // Cek apakah akun admin sudah ada dalam database
    var existingAdmin data.User
    if err := db.Where("role = ?", "admin").First(&existingAdmin).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            fmt.Println("Gagal memeriksa keberadaan akun admin:", err)
            return
        }
    } else {
        fmt.Println("Akun admin sudah ada dalam database. Tidak perlu membuat baru.")
        return
    }

    // Jika akun admin belum ada, buat baru
    // Menghash password sebelum menyimpannya
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Gagal menghash password:", err)
        return
    }

    admin := data.User{
        Role:     "admin",
        Nama:     "admin",
        Email:    "admin@mail.com",
        Password: string(hashedPassword), // Menggunakan password yang di-hash
    }

    if err := db.Create(&admin).Error; err != nil {
        fmt.Println("Seed gagal:", err)
        return
    }

    fmt.Println("Seed berhasil")
}
