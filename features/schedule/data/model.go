package data

import (
	"Medqueue-Alta-BE/features/reservation/data"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	UserID 			uint
	PoliID 			uint
	Hari 			string
	WaktuMulai 		string
	WaktuSelesai 	string
	Kuota 			string
	Reservations 	[]data.Reservation `gorm:"foreign_key:ScheduleID"`
}