package handler

type ScheduleRequest struct {
	PoliKlinik   string `json:"poli" form:"poli"`
	Hari         string `json:"hari" form:"hari"`
	WaktuMulai   string `json:"jam_mulai" form:"jam_mulai"`
	WaktuSelesai string `json:"jam_selesai" form:"jam_selesai"`
	Kuota        string `json:"kuota" form:"kuota"`
}