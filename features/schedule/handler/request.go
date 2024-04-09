package handler

type ScheduleRequest struct {
	PoliID       int    `json:"poli" form:"poli_id"`
	Hari         string `json:"hari" form:"hari"`
	WaktuMulai   string `json:"jam_mulai" form:"jam_mulai"`
	WaktuSelesai string `json:"jam_selesai" form:"jam_selesai"`
	Kuota        int    `json:"kuota" form:"kuota"`
}
