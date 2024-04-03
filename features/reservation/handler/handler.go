package handler

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
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

// handler.go
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

		// Extract adminID from the JWT token
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}
		claims := token.Claims.(jwt.MapClaims)
		adminID := uint(claims["id"].(float64))

		// Create Reservation object
		var inputProcess reservation.Reservation
		inputProcess.PoliKlinik = input.PoliKlinik
		inputProcess.Hari = input.Hari
		inputProcess.WaktuMulai = input.WaktuMulai
		inputProcess.WaktuSelesai = input.WaktuSelesai
		inputProcess.Kuota = input.Kuota

		// Call service to add reservation
		result, err := ct.s.AddReservation(adminID, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		// Return success response
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan reservasi", result))
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
			PoliKlinik:   input.PoliKlinik,
			Hari:         input.Hari,
			WaktuMulai:   input.WaktuMulai,
			WaktuSelesai: input.WaktuSelesai,
			Kuota:        input.Kuota,
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

func (ct *controller) ShowMyReservation() echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		reservation, err := ct.s.GetReservationByOwner(token)
		if err != nil {
			log.Println("gagal mendapat reservasi user:", err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "reservasi pengguna", reservation))
	}
}
