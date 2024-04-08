package services

import (
	"Medqueue-Alta-BE/features/user"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model user.UserModel
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	var registerValidate user.Register
	registerValidate.Nama = newData.Nama
	registerValidate.Email = newData.Email
	registerValidate.Password = newData.Password
	registerValidate.TempatLahir = newData.TempatLahir
	registerValidate.TanggalLahir = newData.TanggalLahir
	registerValidate.JenisKelamin = newData.JenisKelamin
	registerValidate.GolonganDarah = newData.GolonganDarah
	registerValidate.NIK = newData.NIK
	registerValidate.NoBPJS = newData.NoBPJS
	registerValidate.NoTelepon = newData.NoTelepon
	err := s.v.Struct(&registerValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	newPassword, err := s.pm.HashPassword(newData.Password)
	if err != nil {
		return errors.New(helper.ServiceGeneralError)
	}
	newData.Password = newPassword

	err = s.model.AddUser(newData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}
func (s *service) Login(loginData user.User) (user.User, string, error) {
	var loginValidate user.Login
	loginValidate.Email = loginData.Email
	loginValidate.Password = loginData.Password
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return user.User{}, "", err
	}

	dbData, err := s.model.Login(loginValidate.Email)
	if err != nil {
		return user.User{}, "", err
	}

	err = s.pm.ComparePassword(loginValidate.Password, dbData.Password)
	if err != nil {
		return user.User{}, "", errors.New(helper.UserCredentialError)
	}

	token, err := middlewares.GenerateJWT(dbData.ID)
	if err != nil {
		return user.User{}, "", errors.New(helper.ServiceGeneralError)
	}

	return dbData, token, nil
}

func (s *service) Profile(token *jwt.Token) (user.User, error) {
	decodeId := middlewares.DecodeToken(token)
	result, err := s.model.GetUserByID(decodeId)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (s *service) Update(token *jwt.Token, newData user.User) (user.User, error) {
	decodedID := middlewares.DecodeToken(token)

	existingUser, err := s.model.GetUserByID(decodedID)
	if err != nil {
		return user.User{}, errors.New("user not found")
	}

	if newData.Nama != "" {
		existingUser.Nama = newData.Nama
	}

	if newData.Email != "" {
		existingUser.Email = newData.Email
	}

	if newData.Password != "" {
		newPassword, err := s.pm.HashPassword(newData.Password)
		if err != nil {
			return user.User{}, errors.New(helper.ServiceGeneralError)
		}
		existingUser.Password = newPassword
	}

	if newData.TanggalLahir != "" {
		existingUser.TempatLahir = newData.TempatLahir
	}

	if newData.TanggalLahir != "" {
		existingUser.TanggalLahir = newData.TanggalLahir
	}

	if newData.JenisKelamin != "" {
		existingUser.JenisKelamin = newData.JenisKelamin
	}

	if newData.GolonganDarah != "" {
		existingUser.GolonganDarah = newData.GolonganDarah
	}

	if newData.NIK != "" {
		existingUser.NIK = newData.NIK
	}

	if newData.NoBPJS != "" {
		existingUser.NoBPJS = newData.NoBPJS
	}

	if newData.NoTelepon != "" {
		existingUser.NoTelepon = newData.NoTelepon
	}

	result, err := s.model.Update(decodedID, existingUser)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (s *service) Delete(token *jwt.Token) error {
	decodedID := middlewares.DecodeToken(token)
	if decodedID == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.model.Delete(decodedID)
	if err != nil {
		return errors.New("data berhasil dihapus")
	}

	return nil
}
