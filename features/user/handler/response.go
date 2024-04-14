package handler

type LoginResponse struct {
	Role  string `json:"role"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type ProfileResponse struct {
	ID            uint   `json:"id"`
	Role          string `json:"role"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tgl_lahir"`
	JenisKelamin  string `json:"gender"`
	GolonganDarah string `json:"gol_darah"`
	NIK           string `json:"no_nik"`
	NoBPJS        string `json:"no_bpjs"`
	NoTelepon     string `json:"no_telepon"`
}