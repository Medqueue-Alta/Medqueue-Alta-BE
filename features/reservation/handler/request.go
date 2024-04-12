package handler

type ReservationRequest struct {
	PoliID        uint   `json:"poli_id" form:"poli_id"`
	TanggalDaftar string `json:"tanggal_kunjungan" form:"tanggal_kunjungan"`
	ScheduleID    uint   `json:"id_jadwal" form:"id_jadwal"`
	Keluhan       string `json:"keluhan" form:"keluhan"`
	Bpjs          bool   `json:"bpjs" form:"bpjs"`
}
