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

func (rm *model) AddReservation(userid uint, reservasiBaru reservation.Reservation, nama string) (reservation.Reservation, error) {
	var inputProcess = Reservation{PoliID: reservasiBaru.PoliID, TanggalDaftar: reservasiBaru.TanggalDaftar, 
		ScheduleID: reservasiBaru.ScheduleID, Keluhan: reservasiBaru.Keluhan,UserID : userid, Bpjs: reservasiBaru.Bpjs, Status: "waiting", Nama: nama,}
	if err := rm.connection.Create(&inputProcess).Error; err != nil {
		return reservation.Reservation{}, err
	}

	return reservation.Reservation{PoliID: inputProcess.PoliID, TanggalDaftar: inputProcess.TanggalDaftar,
		ScheduleID: inputProcess.ScheduleID, Keluhan: inputProcess.Keluhan, Bpjs: inputProcess.Bpjs, Status: inputProcess.Status, Nama: inputProcess.Nama,}, nil
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

func (rm *model) GetAllReservations() ([]reservation.Reservation, error) {
	var result []reservation.Reservation
	if err := rm.connection.Find(&result).Error; err != nil {
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

func (rm *model) GetReservationByID(reservationID uint) (*reservation.Reservation, error) {
	var result reservation.Reservation
	if err := rm.connection.Where("id = ?", reservationID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (rm *model) GetReservationsByPoliID(poliID uint) ([]reservation.Reservation, error) {
    var result []reservation.Reservation
    if err := rm.connection.Where("poli_id = ?", poliID).Find(&result).Error; err != nil {
        return nil, err
    }
    return result, nil
}