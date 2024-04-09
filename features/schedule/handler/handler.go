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
		inputProcess.PoliID = input.PoliID
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
            PoliID:     input.PoliID,
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


func (ct *controller) ShowAllSchedules() echo.HandlerFunc {
    return func(c echo.Context) error {
        schedule, err := ct.s.GetAllSchedules()
        if err != nil {
            log.Println("gagal mendapat schedule user:", err.Error())
            return c.JSON(http.StatusInternalServerError,
                helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
        }

        return c.JSON(http.StatusOK,
            helper.ResponseFormat(http.StatusOK, "schedule pengguna", schedule))
    }
}

func (ct *controller) ShowScheduleByID() echo.HandlerFunc {
    return func(c echo.Context) error {
        scheduleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
        if err != nil {
            log.Println("error parsing schedule_id:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, "ID schedule tidak valid", nil))
        }

        schedule, err := ct.s.GetScheduleByID(uint(scheduleID))
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


func (ct *controller) ShowSchedulesByPoliID() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Dapatkan nilai parameter query "poli_id"
        poliIDStr := c.QueryParam("poli_id")

        // Konversi poliID dari string ke uint
        poliID, err := strconv.ParseUint(poliIDStr, 10, 64)
        if err != nil {
            log.Println("gagal mengonversi poliID menjadi uint64:", err.Error())
            return c.JSON(http.StatusBadRequest,
                helper.ResponseFormat(http.StatusBadRequest, "Poli ID harus berupa angka", nil))
        }

        // Panggil service untuk mendapatkan jadwal berdasarkan poliID
        schedules, err := ct.s.GetSchedulesByPoliID(uint(poliID))
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


