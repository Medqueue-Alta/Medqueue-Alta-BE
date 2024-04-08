package data

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	UserID        uint
	PoliKlinik    string
	TanggalDaftar string
	Jadwal        string
	Keluhan       string
	Bpjs          bool
}
