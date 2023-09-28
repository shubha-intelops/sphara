package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type AmbulanceService struct {
	ambulanceDao *daos.AmbulanceDao
}

func NewAmbulanceService() (*AmbulanceService, error) {
	ambulanceDao, err := daos.NewAmbulanceDao()
	if err != nil {
		return nil, err
	}
	return &AmbulanceService{
		ambulanceDao: ambulanceDao,
	}, nil
}

func (ambulanceService *AmbulanceService) CreateAmbulance(ambulance *models.Ambulance) (*models.Ambulance, error) {
	return ambulanceService.ambulanceDao.CreateAmbulance(ambulance)
}

func (ambulanceService *AmbulanceService) UpdateAmbulance(id int64, ambulance *models.Ambulance) (*models.Ambulance, error) {
	return ambulanceService.ambulanceDao.UpdateAmbulance(id, ambulance)
}

func (ambulanceService *AmbulanceService) DeleteAmbulance(id int64) error {
	return ambulanceService.ambulanceDao.DeleteAmbulance(id)
}

func (ambulanceService *AmbulanceService) ListAmbulances() ([]*models.Ambulance, error) {
	return ambulanceService.ambulanceDao.ListAmbulances()
}

func (ambulanceService *AmbulanceService) GetAmbulance(id int64) (*models.Ambulance, error) {
	return ambulanceService.ambulanceDao.GetAmbulance(id)
}
