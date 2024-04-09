package handler

type LoginResponse struct {
	Nama  string `json:"nama"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Token string `json:"token"`
}
