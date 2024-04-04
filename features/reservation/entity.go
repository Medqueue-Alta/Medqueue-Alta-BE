package reservation

import (
	"net/http"

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
	AddReservation(adminid uint, reservasiBaru Reservation) (Reservation, error)
	UpdateReservation(adminid uint, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(adminid uint, reservationID uint) error
	GetReservationByOwner(userid uint) ([]Reservation, error)
}

type ReservationService interface {
	AddReservation(reservasiBaru Reservation, req *http.Request) (Reservation, error)
	UpdateReservation(userid *jwt.Token, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid *jwt.Token, reservationID uint) error
	GetReservationByOwner(userid *jwt.Token) ([]Reservation, error)
}

type Reservation struct {
	ID           uint
	AdminID      uint
	PoliKlinik   string
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
	Kuota        int
}
