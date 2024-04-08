package config

import (
	"Medqueue-Alta-BE/features/user/data"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedAdmin(db *gorm.DB) {
    // Menghash password sebelum menyimpannya
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("gagal menghash password:", err)
        return
    }

    admin := data.User{
        Role:     "admin",
        Nama:     "admin",
        Email:    "admin@mail.com",
        Password: string(hashedPassword), // Menggunakan password yang di-hash
    }

    if err := db.Create(&admin).Error; err != nil {
        fmt.Println("seed gagal:", err)
        return
    }

    fmt.Println("seed berhasil")
}
