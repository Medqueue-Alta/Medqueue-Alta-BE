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
	var count int64
    rm.connection.Model(&reservation.Reservation{}).Where("schedule_id = ?", reservasiBaru.ScheduleID).Count(&count)
    reservasiBaru.NoAntrian = count + 1

	var inputProcess = Reservation{PoliID: reservasiBaru.PoliID, TanggalDaftar: reservasiBaru.TanggalDaftar, 
		ScheduleID: reservasiBaru.ScheduleID, Keluhan: reservasiBaru.Keluhan,UserID : userid, Bpjs: reservasiBaru.Bpjs,
		Status: "Waiting", Nama: nama, NoAntrian: reservasiBaru.NoAntrian,}

	if err := rm.connection.Create(&inputProcess).Error; err != nil {
		return reservation.Reservation{}, err
	}

	return reservation.Reservation{PoliID: inputProcess.PoliID, TanggalDaftar: inputProcess.TanggalDaftar,
		ScheduleID: inputProcess.ScheduleID, Keluhan: inputProcess.Keluhan, Bpjs: inputProcess.Bpjs,
		Status: inputProcess.Status, Nama: inputProcess.Nama,NoAntrian: inputProcess.NoAntrian,}, nil
}

func (rm *model) UpdateReservation(reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	var qry = rm.connection.Where(" id = ?", reservationID).Updates(data)
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

    // Membuat peta untuk menyimpan jumlah reservasi terlepas dari status berdasarkan ScheduleID
    counts := make(map[uint]int64)

    // Menghitung jumlah reservasi berdasarkan ScheduleID dari database
    for i := range result {
        // Hitung jumlah reservasi yang sudah "Skipped" atau "Check In" untuk ScheduleID tertentu
        var count int64
        if err := rm.connection.Model(&reservation.Reservation{}).Where("status IN (?) AND schedule_id = ?", []string{"Skipped", "Check In"}, result[i].ScheduleID).Count(&count).Error; err != nil {
            return nil, err
        }
        counts[result[i].ScheduleID] = count
    }

    // Mengatur AntrianNow berdasarkan jumlah reservasi yang sudah Check In atau Skipped
    for i := range result {
        // Set AntrianNow ke total count
        result[i].AntrianNow = counts[result[i].ScheduleID]
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

    // Membuat peta untuk menyimpan jumlah reservasi terlepas dari status berdasarkan ScheduleID
    counts := make(map[uint]int64)

    // Menghitung jumlah reservasi berdasarkan ScheduleID dari database
    for i := range result {
        // Hitung jumlah reservasi yang sudah "Skipped" atau "Check In" untuk ScheduleID tertentu
        var count int64
        if err := rm.connection.Model(&reservation.Reservation{}).Where("status IN (?) AND schedule_id = ?", []string{"Skipped", "Check In"}, result[i].ScheduleID).Count(&count).Error; err != nil {
            return nil, err
        }
        counts[result[i].ScheduleID] = count
    }

    // Mengatur AntrianNow berdasarkan jumlah reservasi yang sudah Check In atau Skipped
    for i := range result {
        // Set AntrianNow ke total count
        result[i].AntrianNow = counts[result[i].ScheduleID]
    }

    return result, nil
}

func (rm *model) GetScheduleByID(scheduleID uint) (*reservation.Schedule, error) {
	var result reservation.Schedule
	if err := rm.connection.Where("id = ?", scheduleID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (rm *model) GetReservationByOwner(userid uint) ([]reservation.Reservation, error) {
    var result []reservation.Reservation
    if err := rm.connection.Where("user_id = ?", userid).Find(&result).Error; err != nil {
        return nil, err
    }

    // Membuat peta untuk menyimpan jumlah reservasi terlepas dari status berdasarkan ScheduleID
    counts := make(map[uint]int64)

    // Menghitung jumlah reservasi berdasarkan ScheduleID dari database
    for i := range result {
        // Hitung jumlah reservasi yang sudah "Skipped" atau "Check In" untuk ScheduleID tertentu
        var count int64
        if err := rm.connection.Model(&reservation.Reservation{}).Where("status IN (?) AND schedule_id = ?", []string{"Skipped", "Check In"}, result[i].ScheduleID).Count(&count).Error; err != nil {
            return nil, err
        }
        counts[result[i].ScheduleID] = count
    }

    // Mengatur AntrianNow berdasarkan jumlah reservasi yang sudah Check In atau Skipped
    for i := range result {
        // Set AntrianNow ke total count
        result[i].AntrianNow = counts[result[i].ScheduleID]
    }

    return result, nil
}

func (rm *model) GetLastReservationByScheduleID(scheduleID uint) (reservation.Reservation, error) {
    var lastReservation reservation.Reservation
    if err := rm.connection.Where("schedule_id = ?", scheduleID).Order("no_antrian desc").First(&lastReservation).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // Return default reservation with NoAntrian = 0
            return reservation.Reservation{NoAntrian: 0}, nil
        }
        // Handle other errors
        return reservation.Reservation{}, err
    }
    return lastReservation, nil
}



