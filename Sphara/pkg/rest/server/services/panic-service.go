package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type PanicService struct {
	panicDao *daos.PanicDao
}

func NewPanicService() (*PanicService, error) {
	panicDao, err := daos.NewPanicDao()
	if err != nil {
		return nil, err
	}
	return &PanicService{
		panicDao: panicDao,
	}, nil
}

func (panicService *PanicService) CreatePanic(panic *models.Panic) (*models.Panic, error) {
	return panicService.panicDao.CreatePanic(panic)
}

func (panicService *PanicService) UpdatePanic(id int64, panic *models.Panic) (*models.Panic, error) {
	return panicService.panicDao.UpdatePanic(id, panic)
}

func (panicService *PanicService) DeletePanic(id int64) error {
	return panicService.panicDao.DeletePanic(id)
}

func (panicService *PanicService) ListPanics() ([]*models.Panic, error) {
	return panicService.panicDao.ListPanics()
}

func (panicService *PanicService) GetPanic(id int64) (*models.Panic, error) {
	return panicService.panicDao.GetPanic(id)
}
