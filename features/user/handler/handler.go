package handler

import (
	"Medqueue-Alta-BE/features/user"
	"Medqueue-Alta-BE/helper"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) user.UserController {
	return &controller{
		service: s,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirimkan tidak sesuai", nil))
		}
		err = ct.service.Register(input)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", nil))
	}
}
func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirimkan tidak sesuai", nil))
		}

		var processData user.User
		processData.Email = input.Email
		processData.Password = input.Password

		result, token, err := ct.service.Login(processData)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		var responseData LoginResponse
		responseData.Email = result.Email
		responseData.Nama = responseData.Nama
		responseData.Role = "Pasien"
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil login", responseData))
	}
}
func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}
		result, err := ct.service.Profile(token)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)

		var inputData user.User
		if err := c.Bind(&inputData); err != nil {
			log.Println("error binding data:", err.Error())
			if strings.Contains(err.Error(), "unsupported") {
				return c.JSON(http.StatusUnauthorized,
					helper.ResponseFormat(http.StatusUnauthorized, "anda tidak bisa mengakses perintah ini", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan pada data input", nil))
		}

		if inputData.Nama == "" && inputData.Email == "" && inputData.Password == "" && inputData.TempatLahir == "" &&
			inputData.TanggalLahir == "" && inputData.JenisKelamin == "" && inputData.GolonganDarah == "" && inputData.NIK == "" &&
			inputData.NoBPJS == "" && inputData.NoTelepon == "" {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan pada data input", nil))
		}

		updatedUser, err := ct.service.Update(token, inputData)
		if err != nil {
			log.Println("failed to update user:", err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, helper.UserInputError, nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mengubah data", updatedUser))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan pada data input", nil))
		}

		err := ct.service.Delete(token)
		if err != nil {
			log.Println("gagal menghapus user:", err.Error())
			return c.JSON(http.StatusUnauthorized,
				helper.ResponseFormat(http.StatusUnauthorized, "anda tidak bisa mengakses perintah ini", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil menghapus data", nil))
	}
}
