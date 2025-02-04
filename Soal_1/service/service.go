package service

import (
	"Soal_1/models"
	"Soal_1/repositories"
	"Soal_1/utils"
	"log"
)

type Soal1Service interface {
	CreateSoal1(soal1 *models.UserData) error
}

type soal1ServiceImpl struct {
	soal1Repo repositories.Soal1Repo
}

func (soal1Service *soal1ServiceImpl) CreateSoal1(soal1 *models.UserData) error {
	if err := soal1.Validate(); err != nil {
		log.Println("Error validasi:", err)
		return &utils.AppError{
			ErrorCode:    101,
			ErrorMessage: err.Error(),
		}
	}

	hashPassword, err := utils.HashPassword(soal1.Password)
	if err != nil {
		log.Println("Error generate password:", err)
		return &utils.AppError{
			ErrorCode:    102,
			ErrorMessage: err.Error(),
		}
	}
	soal1.Password = hashPassword

	return soal1Service.soal1Repo.CreateSoal1(soal1)
}

func NewSoal1Service(soal1Repo repositories.Soal1Repo) Soal1Service {
	return &soal1ServiceImpl{
		soal1Repo: soal1Repo,
	}
}
