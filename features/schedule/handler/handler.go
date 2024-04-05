package handler

import (
	"Medqueue-Alta-BE/features/schedule"
	"Medqueue-Alta-BE/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s schedule.ScheduleService
}

func NewHandler(service schedule.ScheduleService) schedule.ScheduleController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ScheduleRequest
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

		var inputProcess schedule.Schedule
		inputProcess.PoliKlinik = input.PoliKlinik
		inputProcess.Hari = input.Hari
		inputProcess.WaktuMulai = input.WaktuMulai
		inputProcess.WaktuSelesai = input.WaktuSelesai
        inputProcess.Kuota = input.Kuota
		result, err := ct.s.AddSchedule(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan schedule", result))
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

        var input ScheduleRequest
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

        updatedSchedule, err := ct.s.UpdateSchedule(token, uint(id), schedule.Schedule{
            PoliKlinik:     input.PoliKlinik,
			Hari: input.Hari,
			WaktuMulai: input.WaktuMulai,
			WaktuSelesai: input.WaktuSelesai,
            Kuota: input.Kuota,
        })
        if err != nil {
            log.Println("gagal update:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusForbidden, "gagal update", nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "berhasil diperbarui", updatedSchedule))
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

        err = ct.s.DeleteSchedule(token, uint(id))
        if err != nil {
            log.Println("gagal menghapus:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusForbidden, "gagal menghapus", nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "berhasil dihapus", nil))
    }
}


func (ct *controller) ShowMySchedule() echo.HandlerFunc {
    return func(c echo.Context) error {

        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
        }

        schedule, err := ct.s.GetScheduleByOwner(token)
        if err != nil {
            log.Println("gagal mendapat schedule user:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "schedule pengguna", schedule))
    }
}