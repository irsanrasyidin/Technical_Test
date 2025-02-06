package db

import (
	"Soal_3/internal/models"
	"errors"

	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, id int) (*models.UserData, error) {
	var user models.UserData
	result := db.First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, result.Error
}
