package controller

import (
	"Soal_1/config"
	"Soal_1/manager"
	"Soal_1/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.ServiceManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	s.srv.Use(middleware.LoggerMiddleware())

	// controller
	NewSoal1Controller(s.srv, s.usecaseManager.GetSoal1Service())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewServiceManager(repo)

	srv := gin.Default()

	if c.DbConfig.Host == "" || c.AppPort == "" {
		panic("No Host or port define")
	}

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
