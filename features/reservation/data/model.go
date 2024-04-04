package data

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	PoliKlinik   string
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
	Kuota        int
	UserID       uint
	AdminID      uint
}
