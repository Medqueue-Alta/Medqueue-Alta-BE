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
}

type ReservationService interface {
	AddReservation(userid *jwt.Token, reservasiBaru Reservation) (Reservation, error)
	UpdateReservation(userid *jwt.Token, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid *jwt.Token, reservationID uint) error
	GetReservationByOwner(userid *jwt.Token) ([]Reservation, error)
}

type Reservation struct {
	ID 					uint
	UserID 				uint
	PoliKlinik 			string
	TanggalDaftar 		string
	Jadwal 				string
	Keluhan 			string
}