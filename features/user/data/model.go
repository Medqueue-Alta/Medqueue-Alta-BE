package data

import (
	schedule "Medqueue-Alta-BE/features/schedule/data"
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
	Schedule      []schedule.Schedule
}
