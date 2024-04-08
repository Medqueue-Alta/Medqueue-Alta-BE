package handler

type ReservationRequest struct {
	PoliKlinik    string `json:"poli" form:"poli"`
	TanggalDaftar string `json:"tanggal_kunjungan" form:"tanggal_kunjungan"`
	Jadwal        string `json:"jadwal" form:"jadwal"`
	Keluhan       string `json:"keluhan" form:"keluhan"`
}
