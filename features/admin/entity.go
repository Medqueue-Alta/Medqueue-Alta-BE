package admin

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminController interface {
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type AdminService interface {
	Login(loginData Admin) (Admin, string, error)
	Update(token *jwt.Token, newData Admin) (Admin, error)
	Delete(token *jwt.Token) error
}

type AdminModel interface {
	Login(email string) (Admin, error)
	GetUserByID(id uint) (Admin, error)
	Update(id uint, newData Admin) (Admin, error)
	Delete(id uint) error
}

type Admin struct {
	ID       uint   `json:"id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type Login struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,alphanum,min=8"`
}
