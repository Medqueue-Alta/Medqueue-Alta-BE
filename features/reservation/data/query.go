package data

import (
	"Medqueue-Alta-BE/features/reservation"
	"errors"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationModel {
	return &model{
		connection: db,
	}
}

// query.go
func (rm *model) AddReservation(adminID uint, reservasiBaru reservation.Reservation) (reservation.Reservation, error) {
	// Create Reservation object
	var inputProcess = reservation.Reservation{
		PoliKlinik:   reservasiBaru.PoliKlinik,
		Hari:         reservasiBaru.Hari,
		WaktuMulai:   reservasiBaru.WaktuMulai,
		WaktuSelesai: reservasiBaru.WaktuSelesai,
		Kuota:        reservasiBaru.Kuota,
		AdminID:      adminID, // Assuming adminID is the user ID of the admin creating the reservation
	}

	// Insert into the database
	if err := rm.connection.Create(&inputProcess).Error; err != nil {
		return reservation.Reservation{}, err
	}

	// Return the created reservation
	return inputProcess, nil
}

func (s *services) UpdateReservation(adminID *jwt.Token, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	// Get the user ID from the middleware
	userID, err := s.md.GetUserID()
	if err != nil {
		return reservation.Reservation{}, err
	}

	// Check if the user is an admin
	isAdmin, err := s.isAdmin(userID)
	if err != nil {
		return reservation.Reservation{}, err
	}
	if !isAdmin {
		return reservation.Reservation{}, errors.New("only admin can update reservation")
	}

	// Validate input (if necessary)
	// Add your validation logic here if needed

	// Update reservation
	result, err := s.m.UpdateReservation(userID, reservationID, data)
	if err != nil {
		return reservation.Reservation{}, err
	}

	return result, nil
}

func (rm *model) GetReservationByOwner(userid uint) ([]reservation.Reservation, error) {
	var result []reservation.Reservation
	if err := rm.connection.Where("user_id = ?", userid).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (rm *model) DeleteReservation(userid uint, reservationID uint) error {
	result := rm.connection.Unscoped().Where("user_id = ? AND id = ?", userid, reservationID).Delete(&Reservation{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no data affected")
	}

	return nil
}
