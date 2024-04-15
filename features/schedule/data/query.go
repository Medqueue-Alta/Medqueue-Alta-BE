package data

import (
	"Medqueue-Alta-BE/features/schedule"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) schedule.ScheduleModel {
	return &model{
		connection: db,
	}
}

func (rm *model) AddSchedule(userid uint, scheduleBaru schedule.Schedule) (schedule.Schedule, error) {
	var inputProcess = Schedule{PoliID: scheduleBaru.PoliID, Hari: scheduleBaru.Hari, 
		WaktuMulai: scheduleBaru.WaktuMulai, WaktuSelesai: scheduleBaru.WaktuSelesai,
		Kuota: scheduleBaru.Kuota,UserID : userid,}
	if err := rm.connection.Create(&inputProcess).Error; err != nil {
		return schedule.Schedule{}, err
	}

	return schedule.Schedule{PoliID: inputProcess.PoliID, Hari: inputProcess.Hari,
		WaktuMulai: inputProcess.WaktuMulai, WaktuSelesai: inputProcess.WaktuSelesai,
		Kuota: inputProcess.Kuota, Terisi: 0,}, nil
}

func (rm *model) UpdateSchedule(userid uint, scheduleID uint, data schedule.Schedule) (schedule.Schedule, error) {
	var qry = rm.connection.Where("user_id = ? AND id = ?", userid, scheduleID).Updates(data)
	if err := qry.Error; err != nil {
		return schedule.Schedule{}, err
	}

	if qry.RowsAffected < 1 {
		return schedule.Schedule{}, errors.New("no data affected")
	}

	return data, nil
}

func (rm *model) GetAllSchedules() ([]schedule.Schedule, error) {
	var schedules []schedule.Schedule
	if err := rm.connection.Find(&schedules).Error; err != nil {
		return nil, err
	}

	for i, _ := range schedules {
		// Hitung jumlah reservasi untuk jadwal ini
		var count int64
		rm.connection.Model(&schedule.Reservation{}).Where("schedule_id = ?", schedules[i].ID).Count(&count)
		schedules[i].Terisi = count
	}

	return schedules, nil
}


func (rm *model) GetScheduleByID(scheduleID uint) (*schedule.Schedule, error) {
	var result schedule.Schedule
	if err := rm.connection.Where("id = ?", scheduleID).First(&result).Error; err != nil {
		return nil, err
	}

	// Hitung jumlah reservasi yang sudah ada untuk jadwal dengan scheduleID yang sama
	var count int64
	if err := rm.connection.Model(&schedule.Reservation{}).Where("schedule_id = ?", scheduleID).Count(&count).Error; err != nil {
		return nil, err
	}

	// Set nilai Terisi dengan jumlah reservasi yang telah dihitung
	result.Terisi = int64(count)

	return &result, nil
}



func (rm *model) DeleteSchedule(userid uint, scheduleID uint) error {
    result := rm.connection.Unscoped().Where("user_id = ? AND id = ?", userid, scheduleID).Delete(&Schedule{})
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no data affected")
    }

    return nil
}


func (rm *model) GetSchedulesByPoliID(poliID uint) ([]schedule.Schedule, error) {
    var result []schedule.Schedule
    if err := rm.connection.Where("poli_id = ?", poliID).Find(&result).Error; err != nil {
        return nil, err
    }

	for i, _ := range result {
		// Hitung jumlah reservasi untuk jadwal ini
		var count int64
		rm.connection.Model(&schedule.Reservation{}).Where("schedule_id = ?", result[i].ID).Count(&count)
		result[i].Terisi = count
	}
    return result, nil
}