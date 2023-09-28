package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type FireService struct {
	fireDao *daos.FireDao
}

func NewFireService() (*FireService, error) {
	fireDao, err := daos.NewFireDao()
	if err != nil {
		return nil, err
	}
	return &FireService{
		fireDao: fireDao,
	}, nil
}

func (fireService *FireService) CreateFire(fire *models.Fire) (*models.Fire, error) {
	return fireService.fireDao.CreateFire(fire)
}

func (fireService *FireService) UpdateFire(id int64, fire *models.Fire) (*models.Fire, error) {
	return fireService.fireDao.UpdateFire(id, fire)
}

func (fireService *FireService) DeleteFire(id int64) error {
	return fireService.fireDao.DeleteFire(id)
}

func (fireService *FireService) ListFires() ([]*models.Fire, error) {
	return fireService.fireDao.ListFires()
}

func (fireService *FireService) GetFire(id int64) (*models.Fire, error) {
	return fireService.fireDao.GetFire(id)
}
