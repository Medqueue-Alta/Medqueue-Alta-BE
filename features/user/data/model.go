package data

import "Medqueue-BE/features/todo/data"

type User struct {
	Nama     string
	Hp       string `gorm:"type:varchar(13);primaryKey"`
	Password string
	Todos    []data.Todo `gorm:"foreignKey:Pemilik;references:Hp"`
}
