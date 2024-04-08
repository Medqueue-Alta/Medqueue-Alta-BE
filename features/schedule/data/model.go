package data

import (
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	UserID       uint
	PoliID       int
	PoliKlinik   string
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
	Kuota        string
}
