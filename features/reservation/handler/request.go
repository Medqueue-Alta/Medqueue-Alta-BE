package handler

type ReservationRequest struct {
	PoliKlinik    string `json:"poli_klinik" form:"poli_klinik"`
	TanggalDaftar string `json:"tanggal_daftar" form:"tanggal_daftar"`
	Jadwal        string `json:"jadwal" form:"jadwal"`
	Keluhan       string `json:"keluhan" form:"keluhan"`
}
