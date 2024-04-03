package data

import (
	"Medqueue-Alta-BE/features/admin"
	"Medqueue-Alta-BE/features/reservation"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

// AddUser implements user.UserModel.
func (m *model) AddAdmin(newData admin.AdminModel) error {
	panic("unimplemented")
}

func New(db *gorm.DB) admin.AdminModel {
	return &model{
		connection: db,
	}
}

func (m *model) Login(email string) (admin.Admin, error) {
	var result admin.Admin
	if err := m.connection.Where("email = ? ", email).First(&result).Error; err != nil {
		return admin.Admin{}, err
	}
	return result, nil
}

func (m *model) GetUserByID(id uint) (admin.Admin, error) {
	var result admin.Admin
	if err := m.connection.Model(&Admin{}).Where("id = ?", id).First(&result).Error; err != nil {
		return admin.Admin{}, err
	}
	return result, nil
}

func (m *model) Delete(id uint) error {
	result := m.connection.Where("user_id = ?", id).Delete(&reservation.Reservation{})
	if result.Error != nil {
		return result.Error
	}

	result = m.connection.Delete(&Admin{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no data affected")
	}

	return nil
}

func (m *model) Update(id uint, newData admin.Admin) (admin.Admin, error) {
	var updatedAdmin admin.Admin

	tx := m.connection.Begin()

	if newData.Email != "" {
		if err := tx.Model(&admin.Admin{}).Where("id = ?", id).Update("email", newData.Email).Error; err != nil {
			tx.Rollback()
			return admin.Admin{}, err
		}
	}

	if newData.Password != "" {
		if err := tx.Model(&admin.Admin{}).Where("id = ?", id).Update("password", newData.Password).Error; err != nil {
			tx.Rollback()
			return admin.Admin{}, err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return admin.Admin{}, err
	}

	// Ambil data user yang telah diperbarui
	if err := m.connection.First(&updatedAdmin, id).Error; err != nil {
		return admin.Admin{}, err
	}

	return updatedAdmin, nil
}
