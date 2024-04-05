package services

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m reservation.ReservationModel
	v *validator.Validate
}

func NewReservationService(model reservation.ReservationModel) reservation.ReservationService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddReservation(userid *jwt.Token, reservasiBaru reservation.Reservation) (reservation.Reservation, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return reservation.Reservation{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(&reservasiBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return reservation.Reservation{}, err
	}

	result, err := s.m.AddReservation(id, reservasiBaru)
	if err != nil {
		return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *service) UpdateReservation(userid *jwt.Token, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return reservation.Reservation{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(data)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return reservation.Reservation{}, err
	}

	result, err := s.m.UpdateReservation(id, reservationID, data)
	if err != nil {
		return reservation.Reservation{}, errors.New("tidak dapat update")
	}

	return result, nil
}

func (s *service) DeleteReservation(userid *jwt.Token, reservationID uint) error {
    id := middlewares.DecodeToken(userid)
    if id == 0 {
        log.Println("error decode token:", "token tidak ditemukan")
        return errors.New("data tidak valid")
    }

    err := s.m.DeleteReservation(id, reservationID) 
    if err != nil {
        return errors.New("gagal menghapus")
    }

    return nil
}



func (s *service) GetReservationByOwner(userid *jwt.Token) ([]reservation.Reservation, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	reservations, err := s.m.GetReservationByOwner(id)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return reservations, nil
}
