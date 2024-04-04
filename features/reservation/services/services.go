package services

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m  reservation.ReservationModel
	v  *validator.Validate
	md middlewares.JwtInterface
}

func NewReserveService(model reservation.ReservationModel, md middlewares.JwtInterface) reservation.ReservationService {
	return &service{
		m:  model,
		v:  validator.New(),
		md: md,
	}
}

func (s *service) AddReservation(reservasiBaru reservation.Reservation, req *http.Request) (reservation.Reservation, error) {
	// Get the user ID from the middleware
	userID, err := s.md.GetUserID(req)
	if err != nil {
		return reservation.Reservation{}, err
	}

	// Check if the user is an admin
	isAdmin, err := s.isAdmin(userID)
	if err != nil {
		return reservation.Reservation{}, err
	}
	if !isAdmin {
		return reservation.Reservation{}, errors.New("only admin can add reservation")
	}

	// Validate input
	if err := s.validateReservationInput(reservasiBaru); err != nil {
		return reservation.Reservation{}, err
	}

	// Create reservation
	result, err := s.m.AddReservation(userID, reservasiBaru)
	if err != nil {
		return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

// isAdmin checks if the user is an admin
func (s *service) isAdmin(userID uint) (bool, error) {
	// Add logic to check if the user with the given ID is an admin
	// You can query your database or implement any other method to determine admin status
	// For demonstration, we're assuming userID 1 is an admin
	return userID == 1, nil
}

// validateReservationInput validates the reservation input
func (s *service) validateReservationInput(reservation reservation.Reservation) error {
	// Check if Hari is within valid range (Senin-Sabtu)
	validDays := map[string]bool{
		"Senin":  true,
		"Selasa": true,
		"Rabu":   true,
		"Kamis":  true,
		"Jumat":  true,
		"Sabtu":  true,
	}
	if _, ok := validDays[reservation.Hari]; !ok {
		return errors.New("Hari must be one of Senin, Selasa, Rabu, Kamis, Jumat, Sabtu")
	}

	// Add additional validation for WaktuMulai and WaktuSelesai if necessary

	// Check if Kuota is a positive integer
	if reservation.Kuota <= 0 {
		return errors.New("Kuota must be a positive integer")
	}

	return nil
}

func (s *service) UpdateReservation(userid *jwt.Token, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	id := s.md.DecodeToken(userid)
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
	id := s.md.DecodeToken(userid)
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
	id := s.md.DecodeToken(userid)
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
