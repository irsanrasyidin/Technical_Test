package manager

import (
	"Soal_1/config"
	"Soal_1/models"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	GetDB() *gorm.DB
}

type inframanager struct {
	db  *gorm.DB
	cfg config.Config
}

var onceLoadDB sync.Once

func (im *inframanager) GetDB() *gorm.DB {
	onceLoadDB.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", im.cfg.Host, im.cfg.Port, im.cfg.User, im.cfg.Password, im.cfg.Name)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Cannot start app, Error when connect to DB ", err.Error())
		}

		im.db = db

		err = db.AutoMigrate(&models.UserData{})
		if err != nil {
			log.Fatal(err)
		}
	})

	return im.db
}

func (im *inframanager) DBcon() *gorm.DB {
	return im.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := inframanager{
		cfg: config,
	}
	infra.GetDB()
	return &infra
}
