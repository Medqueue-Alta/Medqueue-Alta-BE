package services

import (
	"Medqueue-Alta-BE/features/schedule"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m schedule.ScheduleModel
	v *validator.Validate
}

func NewScheduleService(model schedule.ScheduleModel) schedule.ScheduleService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddSchedule(userid *jwt.Token, scheduleBaru schedule.Schedule) (schedule.Schedule, error) {
	id,role,_ := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return schedule.Schedule{}, errors.New("data tidak valid")
	}

    // Memeriksa peran pengguna
    if role != "admin" {
        log.Println("error: hanya admin yang diizinkan menambah jadwal")
        return schedule.Schedule{}, errors.New("hanya admin yang diizinkan menambah jadwal")
    }

    // Melakukan validasi struktur jadwal baru
	err := s.v.Struct(&scheduleBaru)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return schedule.Schedule{}, err
	}

    // Menambahkan jadwal baru
    result, err := s.m.AddSchedule(id, scheduleBaru)
    if err != nil {
        return schedule.Schedule{}, errors.New(helper.ServerGeneralError)
    }

    return result, nil
}


func (s *service) UpdateSchedule(userid *jwt.Token, scheduleID uint, data schedule.Schedule) (schedule.Schedule, error) {
	id,_,_ := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return schedule.Schedule{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(data)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return schedule.Schedule{}, err
	}

	result, err := s.m.UpdateSchedule(id, scheduleID, data)
	if err != nil {
		return schedule.Schedule{}, errors.New("tidak dapat update")
	}

	return result, nil
}

func (s *service) DeleteSchedule(userid *jwt.Token, scheduleID uint) error {
    id,_,_ := middlewares.DecodeToken(userid)
    if id == 0 {
        log.Println("error decode token:", "token tidak ditemukan")
        return errors.New("data tidak valid")
    }

    err := s.m.DeleteSchedule(id, scheduleID) 
    if err != nil {
        return errors.New("gagal menghapus")
    }

    return nil
}



func (s *service) GetAllSchedules() ([]schedule.Schedule, error) {
	reservations, err := s.m.GetAllSchedules()
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return reservations, nil
}

func (s *service) GetScheduleByID(scheduleID uint) (*schedule.Schedule, error) {
	schedule, err := s.m.GetScheduleByID(scheduleID)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}
	return schedule, nil
}

func (s *service) GetSchedulesByPoliID(poliID uint) ([]schedule.Schedule, error) {
    schedules, err := s.m.GetSchedulesByPoliID(poliID)
    if err != nil {
        return nil, errors.New(helper.ServerGeneralError)
    }
    return schedules, nil
}
