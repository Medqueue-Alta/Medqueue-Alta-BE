package reservation

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ReservationController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ShowAllReservations() echo.HandlerFunc
	ShowReservationByID() echo.HandlerFunc
	ShowMyReservation() echo.HandlerFunc
}

type ReservationModel interface {
	AddReservation(userid uint, reservasiBaru Reservation, nama string) (Reservation, error)
	UpdateReservation(reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid uint, reservationID uint) error
	GetAllReservations() ([]Reservation, error)
	GetReservationByID(reservationID uint) (*Reservation, error)
	GetReservationsByPoliID(poliID uint) ([]Reservation, error)
	GetScheduleByID(scheduleID uint) (*Schedule, error)
	GetReservationByOwner(userid uint) ([]Reservation, error)
}

type ReservationService interface {
	AddReservation(userid *jwt.Token, reservasiBaru Reservation) (Reservation, error)
	UpdateReservation(userid *jwt.Token, reservationID uint, data Reservation) (Reservation, error)
	DeleteReservation(userid *jwt.Token, reservationID uint) error
	GetAllReservations() ([]Reservation, error)
	GetReservationByID(reservationsID uint) (*Reservation, error)
	GetReservationsByPoliID(poliID uint) ([]Reservation, error)
	GetReservationByOwner(userid *jwt.Token) ([]Reservation, error)
}

type Reservation struct {
	ID 					uint   `gorm:"auto_increment" json:"reservations_id"`
	UserID				uint   `json:"user_id"`
	Nama				string `json:"nama"`
	ScheduleID			uint   `json:"id_jadwal"`
	PoliID 			    uint   `json:"poli_id"`
	TanggalDaftar 		string `form:"tanggal_kunjungan" json:"tanggal_kunjungan"`
	Keluhan 			string `json:"keluhan"`
	Bpjs 				bool   `json:"bpjs"`
	Status				string `gorm:"default:waiting" json:"status"`
	NoAntrian			int64   `json:"antrian_anda"`
	AntrianNow          int64   `json:"antrian_sekarang" gorm:"default:1"`
}

type Schedule struct {
	ID 					uint   `json:"schedule_id"`
	PoliID 				uint   `json:"poli_id"`
	Hari 				string `json:"hari"`
	WaktuMulai 			string `json:"jam_mulai"`
	WaktuSelesai		string `json:"jam_selesai"`
	Kuota	 			uint   `json:"kuota"`
}
