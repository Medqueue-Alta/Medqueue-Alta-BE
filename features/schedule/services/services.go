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
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return schedule.Schedule{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(&scheduleBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return schedule.Schedule{}, err
	}

	result, err := s.m.AddSchedule(id, scheduleBaru)
	if err != nil {
		return schedule.Schedule{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *service) UpdateSchedule(userid *jwt.Token, scheduleID uint, data schedule.Schedule) (schedule.Schedule, error) {
	id := middlewares.DecodeToken(userid)
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
    id := middlewares.DecodeToken(userid)
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



func (s *service) GetScheduleByOwner(userid *jwt.Token) ([]schedule.Schedule, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	reservations, err := s.m.GetScheduleByOwner(id)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return reservations, nil
}
