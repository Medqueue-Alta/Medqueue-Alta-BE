package handler

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s reservation.ReservationService
}

func NewHandler(service reservation.ReservationService) reservation.ReservationController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ReservationRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

        id, nama, _ := middlewares.DecodeToken(token)

		var inputProcess reservation.Reservation
		inputProcess.PoliID = input.PoliID
		inputProcess.TanggalDaftar = input.TanggalDaftar
		inputProcess.ScheduleID = input.ScheduleID
		inputProcess.Keluhan = input.Keluhan
        inputProcess.Bpjs = input.Bpjs
        inputProcess.Nama = nama
        inputProcess.Status = "Waiting"
        
		result, err := ct.s.AddReservation(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}


        var responseData ReservationResponse
        responseData.ID = result.ID
        responseData.UserID = id
        responseData.Nama = nama
        responseData.ScheduleID = result.ScheduleID
        responseData.PoliID = result.PoliID
        responseData.TanggalDaftar = result.TanggalDaftar
        responseData.Keluhan = result.Keluhan
        responseData.Bpjs = result.Bpjs
        responseData.Status = result.Status
        responseData.NoAntrian = result.NoAntrian

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan reservasi", responseData))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
    return func(c echo.Context) error {
        idStr := c.Param("id")
        id, err := strconv.ParseUint(idStr, 10, 64)
        if err != nil {
            log.Println("error parsing ID:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        var input ReservationRequest
        if err := c.Bind(&input); err != nil {
            log.Println("error bind data:", err.Error())
            if strings.Contains(err.Error(), "unsupported") {
                return c.JSON(http.StatusUnsupportedMediaType,
                    helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
            }
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        updatedReservation, err := ct.s.UpdateReservation(token, uint(id), reservation.Reservation{
            PoliID:     input.PoliID,
			TanggalDaftar: input.TanggalDaftar,
			ScheduleID: input.ScheduleID,
			Keluhan: input.Keluhan,
            Bpjs: input.Bpjs,
            Status: input.Status,
        })
        if err != nil {
            log.Println("gagal update:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusForbidden, "gagal update", nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "berhasil diperbarui", updatedReservation))
    }
}

func (ct *controller) Delete() echo.HandlerFunc {
    return func(c echo.Context) error {
        idStr := c.Param("id")
        id, err := strconv.ParseUint(idStr, 10, 64)
        if err != nil {
            log.Println("error parsing ID:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        err = ct.s.DeleteReservation(token, uint(id))
        if err != nil {
            log.Println("gagal menghapus:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusForbidden, "gagal menghapus", nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "berhasil dihapus", nil))
    }
}

func (ct *controller) ShowReservationByID() echo.HandlerFunc {
    return func(c echo.Context) error {
        reservationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
        if err != nil {
            log.Println("error parsing schedule_id:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, "ID schedule tidak valid", nil))
        }

        schedule, err := ct.s.GetReservationByID(uint(reservationID))
        if err != nil {
            log.Println("error get post by ID:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        if schedule == nil {
            return c.JSON(http.StatusNotFound,
                helper.ResponseFormat(http.StatusNotFound, "Schedule tidak ditemukan", nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "Schedule", schedule))
    }
}


func (ct *controller) ShowAllReservations() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Dapatkan nilai parameter query "poli_id"
        poliIDStr := c.QueryParam("poli_id")

        if poliIDStr == "" {
            schedules, err := ct.s.GetAllReservations()
            if err != nil {
                log.Println("gagal mendapatkan semua jadwal:", err.Error())
                return c.JSON(http.StatusInternalServerError,
                    helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
            }
            return c.JSON(http.StatusOK,
                helper.ResponseFormat(http.StatusOK, "Semua Jadwal", schedules))
        }

        // Konversi poliID dari string ke uint
        poliID, err := strconv.ParseUint(poliIDStr, 10, 64)
        if err != nil {
            log.Println("gagal mengonversi poliID menjadi uint64:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, "Poli ID harus berupa angka", nil))
        }

        // Panggil service untuk mendapatkan jadwal berdasarkan poliID
        schedules, err := ct.s.GetReservationsByPoliID(uint(poliID))
        if err != nil {
            log.Println("gagal mendapat jadwal untuk poliID", poliID, ":", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        // Kembalikan jadwal yang sesuai dalam respons
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "Jadwal untuk poliID "+strconv.FormatUint(poliID, 10), schedules))
    }
}

func (ct *controller) ShowMyReservation() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan token pengguna dari context
        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        // Memanggil service untuk mendapatkan aktivitas pengguna
        reservation, err := ct.s.GetReservationByOwner(token)
        if err != nil {
            log.Println("error get user reservation:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        // Mengembalikan respons HTTP dengan aktivitas pengguna
        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "reservasi pengguna", reservation))
    }
}