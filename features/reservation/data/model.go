package data

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	UserID 				uint
	ScheduleID			uint
	PoliID 			    uint
	TanggalDaftar 		string
	Keluhan 			string
	Bpjs				bool
	Status				string `gorm:"default:waiting"`
}
