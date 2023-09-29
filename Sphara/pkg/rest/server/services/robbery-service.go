package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type RobberyService struct {
	robberyDao *daos.RobberyDao
}

func NewRobberyService() (*RobberyService, error) {
	robberyDao, err := daos.NewRobberyDao()
	if err != nil {
		return nil, err
	}
	return &RobberyService{
		robberyDao: robberyDao,
	}, nil
}

func (robberyService *RobberyService) CreateRobbery(robbery *models.Robbery) (*models.Robbery, error) {
	return robberyService.robberyDao.CreateRobbery(robbery)
}

func (robberyService *RobberyService) UpdateRobbery(id int64, robbery *models.Robbery) (*models.Robbery, error) {
	return robberyService.robberyDao.UpdateRobbery(id, robbery)
}

func (robberyService *RobberyService) DeleteRobbery(id int64) error {
	return robberyService.robberyDao.DeleteRobbery(id)
}

func (robberyService *RobberyService) ListRobberies() ([]*models.Robbery, error) {
	return robberyService.robberyDao.ListRobberies()
}

func (robberyService *RobberyService) GetRobbery(id int64) (*models.Robbery, error) {
	return robberyService.robberyDao.GetRobbery(id)
}
