package data

import (
	"Medqueue-Alta-BE/features/reservation/data"
)

type Admin struct {
	ID           uint `gorm:"primary_key;auto_increment"`
	Email        string
	Password     string
	Reservations []data.Reservation `gorm:"foreign_key:UserID"`
}
