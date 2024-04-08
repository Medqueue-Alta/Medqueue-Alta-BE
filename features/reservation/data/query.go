package data

import (
	"Medqueue-Alta-BE/features/reservation"
	"errors"

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

func (rm *model) AddReservation(userid uint, reservasiBaru reservation.Reservation) (reservation.Reservation, error) {
	var inputProcess = Reservation{PoliKlinik: reservasiBaru.PoliKlinik, TanggalDaftar: reservasiBaru.TanggalDaftar,
		Jadwal: reservasiBaru.Jadwal, Keluhan: reservasiBaru.Keluhan, UserID: userid}
	if err := rm.connection.Create(&inputProcess).Error; err != nil {
		return reservation.Reservation{}, err
	}

	return reservation.Reservation{PoliKlinik: inputProcess.PoliKlinik, TanggalDaftar: inputProcess.TanggalDaftar,
		Jadwal: inputProcess.Jadwal, Keluhan: inputProcess.Keluhan}, nil
}

func (rm *model) UpdateReservation(userid uint, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	var qry = rm.connection.Where("user_id = ? AND id = ?", userid, reservationID).Updates(data)
	if err := qry.Error; err != nil {
		return reservation.Reservation{}, err
	}

	if qry.RowsAffected < 1 {
		return reservation.Reservation{}, errors.New("no data affected")
	}

	return data, nil
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

func (rm *model) GetUserByID(userID uint) (reservation.User, error) {
	var user reservation.User
	if err := rm.connection.First(&user, userID).Error; err != nil {
		return reservation.User{}, err
	}
	return user, nil
}
