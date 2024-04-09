package reservation

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ReservationController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ShowMyReservation() echo.HandlerFunc
}

type ReservationModel interface {
	AddReservation(userid uint, reservasiBaru Reservation) (Reservation, error)
	UpdateReservation(userid uint, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid uint, reservationID uint) error
	GetReservationByOwner(userid uint) ([]Reservation, error)
	GetUserByID(userID uint) (User, error)
}

type ReservationService interface {
	AddReservation(userid *jwt.Token, reservasiBaru Reservation) (Reservation, error)
	UpdateReservation(userid *jwt.Token, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid *jwt.Token, reservationID uint) error
	GetReservationByOwner(userid *jwt.Token) ([]Reservation, error)
}

type Reservation struct {
	ID 					uint   `json:"reservations_id"`
	ScheduleID			uint   `json:"id_jadwal"`
	PoliID 			    uint   `json:"poli_id"`
	TanggalDaftar 		string `form:"tanggal_kunjungan" json:"tanggal_kunjungan"`
	Keluhan 			string `json:"keluhan"`
	Bpjs 				bool   `json:"bpjs"`
	Status				string `gorm:"default:waiting" json:"status"`
}

type User struct {
	ID 				uint		`json:"id"`
	Role			string		`json:"role"`
	Nama 			string		`json:"nama"`
	Email 			string		`json:"email"`
}