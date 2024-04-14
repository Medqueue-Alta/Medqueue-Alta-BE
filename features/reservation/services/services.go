package services

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"errors"
	"log"
	"time"

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
    id, nama, role := middlewares.DecodeToken(userid)
    if id == 0 {
        log.Println("error decode token:", "token tidak ditemukan")
        return reservation.Reservation{}, errors.New("data tidak valid")
    }

    if role != "pasien" {
        log.Println("error: hanya pasien yang diizinkan menambah reservasi")
        return reservation.Reservation{}, errors.New("hanya pasien yang diizinkan menambah reservasi")
    }

    // Dapatkan jadwal terkait dengan reservasi
    schedule, err := s.m.GetScheduleByID(reservasiBaru.ScheduleID)
    if err != nil {
        log.Println("error:", err.Error())
        return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
    }

    // Periksa apakah poli_id pada jadwal sama dengan poli_id pada reservasi
    if schedule.PoliID != reservasiBaru.PoliID {
        log.Println("error: poli_id pada reservasi tidak sesuai dengan jadwal")
        return reservation.Reservation{}, errors.New("poli_id pada reservasi tidak sesuai dengan jadwal")
    }


    reservasiTime, err := time.Parse("02-01-2006", reservasiBaru.TanggalDaftar)
    if err != nil {
        log.Println("error:", err.Error())
        return reservation.Reservation{}, err
    }

    if schedule.Hari != reservasiTime.Weekday().String() {
        return reservation.Reservation{}, errors.New("tanggal reservasi tidak sesuai dengan jadwal")
    }
    log.Println(reservasiBaru.TanggalDaftar)

    // Lakukan operasi tambah reservasi jika poli_id sesuai
    result, err := s.m.AddReservation(id, reservasiBaru, nama)
    if err != nil {
        return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
    }


    return result, nil
}





func (s *service) UpdateReservation(userid *jwt.Token, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
    id, _, role := middlewares.DecodeToken(userid)
    if id == 0 {
        log.Println("error decode token:", "token tidak ditemukan")
        return reservation.Reservation{}, errors.New("data tidak valid")
    }

    // Jika pengguna bukan admin, pastikan mereka hanya dapat memperbarui reservasi yang mereka buat sendiri
    if role != "admin" {
        // Lakukan pengecekan apakah pengguna yang memperbarui reservasi adalah pemiliknya
        reserv, err := s.m.GetReservationByID(reservationID)
        if err != nil {
            log.Println("error:", err.Error())
            return reservation.Reservation{}, errors.New("tidak dapat menemukan reservasi")
        }

        if reserv.UserID != id {
            log.Println("error: hanya pemilik reservasi yang diizinkan melakukan update")
            return reservation.Reservation{}, errors.New("hanya pemilik reservasi yang diizinkan melakukan update")
        }
    }

    err := s.v.Struct(data)
    if err != nil {
        log.Println("error validasi aktivitas", err.Error())
        return reservation.Reservation{}, err
    }

    // Menggunakan ID pengguna (userID) dan ID reservasi (reservationID) dalam pembaruan
    result, err := s.m.UpdateReservation(reservationID, data)
    if err != nil {
        return reservation.Reservation{}, errors.New("tidak dapat update 1")
    }

    return result, nil
}





func (s *service) DeleteReservation(userid *jwt.Token, reservationID uint) error {
    id,_,_ := middlewares.DecodeToken(userid)
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



func (s *service) GetAllReservations() ([]reservation.Reservation, error) {
	reservations, err := s.m.GetAllReservations()
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return reservations, nil
}

func (s *service) GetReservationByID(reservationID uint) (*reservation.Reservation, error) {
	schedule, err := s.m.GetReservationByID(reservationID)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}
	return schedule, nil
}

func (s *service) GetReservationsByPoliID(poliID uint) ([]reservation.Reservation, error) {
    schedules, err := s.m.GetReservationsByPoliID(poliID)
    if err != nil {
        return nil, errors.New(helper.ServerGeneralError)
    }
    return schedules, nil
}