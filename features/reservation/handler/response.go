package handler

type ReservationResponse struct {
	ID            uint   `gorm:"auto_increment" json:"reservations_id"`
	UserID        uint   `json:"user_id"`
	Nama          string `json:"nama"`
	ScheduleID    uint   `json:"id_jadwal"`
	PoliID        uint   `json:"poli_id"`
	TanggalDaftar string `form:"tanggal_kunjungan" json:"tanggal_kunjungan"`
	Keluhan       string `json:"keluhan"`
	Bpjs          bool   `json:"bpjs"`
	Status        string `gorm:"default:waiting" json:"status"`
	NoAntrian     int64  `json:"antrian_anda"`
	AntrianNow    int64  `json:"antrian_sekarang" gorm:"default:1"`
}