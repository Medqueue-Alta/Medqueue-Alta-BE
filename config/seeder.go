package config

import (
	"Medqueue-Alta-BE/features/user/data"
	"fmt"

	"gorm.io/gorm"
)

func seedAdmin(db *gorm.DB) {
	admin:= data.User{
		Role: "admin",
		Nama: "admin",
		Email: "admin@mail.com",
		Password: "admin123",
	}

	if err := db.Create(&admin).Error; err != nil {
		fmt.Println("seed gagal", err)
		return
	}

	fmt.Println("seed berhasil")
}