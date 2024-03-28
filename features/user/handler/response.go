package handler

type LoginResponse struct {
	Email string `json:"email"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}

type ProfileResponse struct {
	UserID    uint
	Nama      string
	Email     string
	Password  string
	Tgl_Lahir string
	Bpjs      string
	Nik       string
	Darah     string
	Telp      int
	Gender    bool
}
