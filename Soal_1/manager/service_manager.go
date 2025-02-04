package manager

import (
	"Soal_1/service"
	"sync"
)

type ServiceManager interface {
	GetSoal1Service() service.Soal1Service
}

type servicemanager struct {
	repoManager RepoManager

	soal1Service service.Soal1Service
}

var onceLoadSoal1Service sync.Once

func (um *servicemanager) GetSoal1Service() service.Soal1Service{
	onceLoadSoal1Service.Do(func() {
		um.soal1Service = service.NewSoal1Service(um.repoManager.GetSoal1Repo())
	})
	return um.soal1Service
}

func NewServiceManager(repoManager RepoManager) ServiceManager {
	return &servicemanager{
		repoManager: repoManager,
	}
}
