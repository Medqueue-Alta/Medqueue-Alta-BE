package handler

type ReservationRequest struct {
	PoliKlinik   string `json:"poli_klinik" form:"poli_klinik"`
	Hari         string `json:"hari" form:"hari"`
	WaktuMulai   string `json:"waktumulai" form:"waktumulai"`
	WaktuSelesai string `json:"waktuselesai" form:"waktuselesai"`
	Kuota        int    `json:"kuota" form:"kuota"`
}
