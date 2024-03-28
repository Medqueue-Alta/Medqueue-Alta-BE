package data

import (
	"Medqueue-BE/features/user"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

// yang kita butuhkan adalah sebuah model
// tapi kenapa return function kok bukan obyek model?

func New(db *gorm.DB) user.UserModel {
	return &model{
		connection: db,
	}
}

func (m *model) InsertUser(newData user.User) error {
	err := m.connection.Create(&newData).Error
	// if err != nil {
	// 	return err
	// }

	if err != nil {
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		log.Println("error database process:", err)

		// 	}
		// }()
		return errors.New("terjadi masalah pada database")
	}

	return nil
}

func (m *model) cekUser(email string) bool {
	var data User
	if err := m.connection.Where("email = ?", email).First(&data).Error; err != nil {
		return false
	}
	return true
}

func (m *model) UpdateUser(email string, data user.User) error {
	if err := m.connection.Model(&data).Where("email = ?", email).Update("nama", data.Nama).Update("password", data.Password).Error; err != nil {
		return err
	}
	return nil
}

func (m *model) GetAllUser() ([]user.User, error) {
	var result []user.User

	if err := m.connection.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (m *model) GetUserByEmail(email string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("email = ?", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) Login(email string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("email = ? ", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) GetLastUserID() (uint, error) {
	var lastUser User

	// query untuk mendapatkan userID terakhir berdasarkan id terbesar
	if err := m.connection.Order("user_id desc").First(&lastUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// tabel kosong, return 0 sebagai userID pertama
			return 0, nil
		}
		return 0, err
	}

	return lastUser.userID, nil
}
