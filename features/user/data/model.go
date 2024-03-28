package data

import "Medqueue-BE/features/todo/data"

type User struct {
	userID    uint
	Nama      string
	Email     string `gorm:"type:varchar(13);primaryKey"`
	Password  string
	Tgl_Lahir string
	Bpjs      string
	Nik       string
	Darah     string
	Telp      int
	gender    bool
	Todos     []data.Todo `gorm:"foreignKey:Pemilik;references:Hp"`
}
