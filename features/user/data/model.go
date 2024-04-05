package data

import (
	reservation "Medqueue-Alta-BE/features/reservation/data"
	schedule "Medqueue-Alta-BE/features/schedule/data"
)

type User struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	Role		  string `gorm:"default:pasien"`
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
	Reservations  []reservation.Reservation  `gorm:"foreign_key:UserID"`
	Schedules  	  []schedule.Schedule  `gorm:"foreign_key:UserID"`
}
