package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// bagian yang berisi KONTRAK mengenai obyek yang digunakan / disepakati dalam proses coding kalian

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
}

type UserModel interface {
	InsertUser(newData User) error
	UpdateUser(email string, data User) error
	Login(email string) (User, error)
	GetUserByEmail(email string) (User, error)
}

type User struct {
	Nama     string
	Email    string
	Password string
}

type Login struct {
	Email    string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}

type Register struct {
	Nama     string `validate:"required,alpha"`
	Email    string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}
