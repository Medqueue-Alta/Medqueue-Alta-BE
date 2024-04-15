package schedule

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ScheduleController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ShowScheduleByID() echo.HandlerFunc
	ShowAllSchedules() echo.HandlerFunc
}

type ScheduleModel interface {
	AddSchedule(userid uint, scheduleBaru Schedule) (Schedule, error)
	UpdateSchedule(userid uint, scheduleID uint, data Schedule) (Schedule, error)
	DeleteSchedule(userid uint, scheduleID uint) error
	GetAllSchedules() ([]Schedule, error)
	GetScheduleByID(scheduleID uint) (*Schedule, error)
	GetSchedulesByPoliID(poliID uint) ([]Schedule, error)
}

type ScheduleService interface {
	AddSchedule(userid *jwt.Token, scheduleBaru Schedule) (Schedule, error)
	UpdateSchedule(userid *jwt.Token, scheduleID uint, data Schedule) (Schedule, error)
	DeleteSchedule(userid *jwt.Token, scheduleID uint) error
	GetAllSchedules() ([]Schedule, error)
	GetScheduleByID(scheduleID uint) (*Schedule, error)
	GetSchedulesByPoliID(poliID uint) ([]Schedule, error)
}

type Schedule struct {
	ID 					uint   `json:"schedule_id"`
	PoliID 				uint   `json:"poli_id"`
	Hari 				string `json:"hari"`
	WaktuMulai 			string `json:"jam_mulai"`
	WaktuSelesai		string `json:"jam_selesai"`
	Kuota	 			uint   `json:"kuota"`
	Terisi 	 			int64  `json:"terisi"`
}

type Reservation struct {
	ID 					uint   `gorm:"auto_increment" json:"reservations_id"`
	UserID				uint   `json:"user_id"`
	Nama				string `json:"nama"`
	ScheduleID			uint   `json:"id_jadwal"`
	NoAntrian			int64   `json:"no_antrian"`
}
