package handler

type LoginResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
