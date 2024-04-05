package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc 
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
	Update(token *jwt.Token, newData User) (User, error)
	Delete(token *jwt.Token) error
}

type UserModel interface {
	AddUser(newData User) error
	Login(email string) (User, error)
	GetUserByID(id uint) (User, error)
	Update(id uint, newData User) (User, error)
	Delete(id uint) error
}

type User struct {
	ID 				uint		`json:"id"`
	Role			string		`json:"role"`
	Nama 			string		`form:"nama" json:"nama"`
	Email 			string		`form:"email" json:"email"`
	Password 		string		`form:"password" json:"password"`
	TempatLahir 	string		`form:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir 	string		`form:"tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin 	string		`form:"jenis_kelamin" json:"jenis_kelamin"`
	GolonganDarah 	string		`form:"golongan_darah" json:"golongan_darah"`
	NIK 			string		`form:"nik" json:"nik"`
	NoBPJS 			string		`form:"no_bpjs" json:"no_bpjs"`
	NoTelepon 		string		`form:"no_telepon" json:"no_telepon"`
}

type Login struct {
	Email 		string `json:"email" form:"email" validate:"required"`
	Password 	string `json:"password" form:"password" validate:"required,alphanum,min=8"`
}

type Register struct {
	Nama 			string		`form:"nama" json:"nama" validate:"required"`
	Email 			string		`form:"email" json:"email" validate:"required"`
	Password 		string		`form:"password" json:"password" validate:"required"`
	TempatLahir 	string		`form:"tempat_lahir" json:"tempat_lahir" validate:"required"`
	TanggalLahir 	string		`form:"tanggal_lahir" json:"tanggal_lahir" validate:"required"`
	JenisKelamin 	string		`form:"jenis_kelamin" json:"jenis_kelamin" validate:"required"`
	GolonganDarah 	string		`form:"golongan_darah" json:"golongan_darah" validate:"required"`
	NIK 			string		`form:"nik" json:"nik" validate:"required"`
	NoBPJS 			string		`form:"no_bpjs" json:"no_bpjs" validate:"required"`
	NoTelepon 		string		`form:"no_telepon" json:"no_telepon" validate:"required"`
}