package repositories

import (
	"Soal_1/models"
	"Soal_1/utils"
	"strings"

	"gorm.io/gorm"
)

type Soal1Repo interface {
	CreateSoal1(soal1l *models.UserData) error
}

type soal1RepoImpl struct {
	db *gorm.DB
}

func (soal1lRepo *soal1RepoImpl) CreateSoal1(soal1l *models.UserData) error {
	if err := soal1lRepo.db.Create(&soal1l).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return &utils.AppError{
				ErrorCode:    201, 
				ErrorMessage: "Email sudah terdaftar.",
			}
		}
		return &utils.AppError{
			ErrorCode:    202,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}
func Newsoal1Repo(db *gorm.DB) Soal1Repo {
	return &soal1RepoImpl{
		db: db,
	}
}
