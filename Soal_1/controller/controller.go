package controller

import (
	"Soal_1/models"
	"Soal_1/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Soal1Controller struct {
	soal1Service service.Soal1Service
}

func (soal1Controller Soal1Controller) CreateSoal1(ctx *gin.Context) {
	var soal1 models.UserData

	if err := ctx.ShouldBindJSON(&soal1); err != nil {
		log.Println("Error parsing request:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": err.Error(),
		})
		return
	}

	err := soal1Controller.soal1Service.CreateSoal1(&soal1)
	if err != nil {
		log.Println("Error creating user:", err)
		if strings.Contains(err.Error(), "Email sudah terdaftar") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": err.Error(),
			})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": err.Error(),
			})
		}		
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Berhasil membuat user dengan ID = " + strconv.Itoa(int(soal1.ID)) + ", Nama = " + soal1.Name + ", Email = " + soal1.Email,
	})
}

func NewSoal1Controller(srv *gin.Engine, soal1Service service.Soal1Service) *Soal1Controller {
	soal1Controller := &Soal1Controller{
		soal1Service: soal1Service,
	}

	srv.POST("/users", soal1Controller.CreateSoal1)

	return soal1Controller
}
