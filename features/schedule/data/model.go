package data

import (
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	UserID       uint
	PoliID       int
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
	Kuota        int
}
