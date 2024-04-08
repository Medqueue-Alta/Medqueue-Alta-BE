package services

import (
	"Medqueue-Alta-BE/features/reservation"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"context"
	"errors"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/option"
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

func (s *service) UpdateReservation(userid *jwt.Token, reservationID uint, data reservation.Reservation) (reservation.Reservation, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return reservation.Reservation{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(data)
	if err != nil {
		log.Println("error validasi aktivitas", err.Error())
		return reservation.Reservation{}, err
	}

	result, err := s.m.UpdateReservation(id, reservationID, data)
	if err != nil {
		return reservation.Reservation{}, errors.New("tidak dapat update")
	}

	return result, nil
}

func (s *service) DeleteReservation(userid *jwt.Token, reservationID uint) error {
	id := middlewares.DecodeToken(userid)
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

func (s *service) GetReservationByOwner(userid *jwt.Token) ([]reservation.Reservation, error) {
	id := middlewares.DecodeToken(userid)
	if id == 0 {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	reservations, err := s.m.GetReservationByOwner(id)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return reservations, nil
}

func SendNotificationToAdmin(ctx context.Context, adminRegistrationToken string, messageTitle string, messageBody string) error {
	opt := option.WithCredentialsFile("path/to/your/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Get a client
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// Construct the message
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: messageTitle,
			Body:  messageBody,
		},
		Token: adminRegistrationToken,
	}

	// Send the message
	_, err = client.Send(ctx, message)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
		return err
	}

	return nil
}

func (s *service) AddReservation(userid *jwt.Token, reservasiBaru reservation.Reservation) (reservation.Reservation, error) {
	ctx := context.Background()

	userInfo, err := middlewares.DecodeTokenWithClaims(userid)
	if err != nil {
		log.Println("error decode token", err.Error())
		return reservation.Reservation{}, err
	}

	user, err := s.m.GetUserByID(userInfo)
	if err != nil {
		log.Println("error mendapatkan informasi pengguna", err.Error())
		return reservation.Reservation{}, err
	}

	if user.Role != "pasien" {
		log.Println("error: hanya pasien yang diizinkan menambah reservasi")
		return reservation.Reservation{}, errors.New("hanya admin yang diizinkan menambah reservasi")
	}

	err = s.v.Struct(&reservasiBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return reservation.Reservation{}, err
	}

	result, err := s.m.AddReservation(user.ID, reservasiBaru)
	if err != nil {
		return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
	}

	// Sending notification to admin
	adminRegistrationToken := "admin_registration_token_here"
	messageTitle := "New Reservation Request"
	messageBody := "A new reservation request has been made. Please review it."
	err = SendNotificationToAdmin(ctx, adminRegistrationToken, messageTitle, messageBody)
	if err != nil {
		log.Println("error sending notification to admin:", err.Error())
		// Handle error sending notification
	}

	return result, nil
}
