package data

import (
	"Medqueue-Alta-BE/features/reservation/data"
)

type User struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	Nama          string
	Email         string
	Password      string
	TempatLahir   string
	TanggalLahir  string
	JenisKelamin  string
	GolonganDarah string
	NIK           string
	NoBPJS        string
	NoTelepon     string
	Reservations  []data.Reservation  `gorm:"foreign_key:UserID"`
}
