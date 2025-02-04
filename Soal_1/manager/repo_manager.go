package manager

import (
	"Soal_1/repositories"
	"sync"
)

type RepoManager interface {
	GetSoal1Repo() repositories.Soal1Repo
}

type repomanager struct {
	infraManager InfraManager

	soal1Repo repositories.Soal1Repo
}

var onceLoadSoal1Repo sync.Once

func (rm *repomanager) GetSoal1Repo() repositories.Soal1Repo {
	onceLoadSoal1Repo.Do(func() {
		rm.soal1Repo = repositories.Newsoal1Repo(rm.infraManager.GetDB())
	})

	return rm.soal1Repo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repomanager{
		infraManager: infraManager,
	}
}
