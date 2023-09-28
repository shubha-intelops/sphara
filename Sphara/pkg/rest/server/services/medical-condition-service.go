package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type MedicalConditionService struct {
	medicalConditionDao *daos.MedicalConditionDao
}

func NewMedicalConditionService() (*MedicalConditionService, error) {
	medicalConditionDao, err := daos.NewMedicalConditionDao()
	if err != nil {
		return nil, err
	}
	return &MedicalConditionService{
		medicalConditionDao: medicalConditionDao,
	}, nil
}

func (medicalConditionService *MedicalConditionService) CreateMedicalCondition(medicalCondition *models.MedicalCondition) (*models.MedicalCondition, error) {
	return medicalConditionService.medicalConditionDao.CreateMedicalCondition(medicalCondition)
}

func (medicalConditionService *MedicalConditionService) UpdateMedicalCondition(id int64, medicalCondition *models.MedicalCondition) (*models.MedicalCondition, error) {
	return medicalConditionService.medicalConditionDao.UpdateMedicalCondition(id, medicalCondition)
}

func (medicalConditionService *MedicalConditionService) DeleteMedicalCondition(id int64) error {
	return medicalConditionService.medicalConditionDao.DeleteMedicalCondition(id)
}

func (medicalConditionService *MedicalConditionService) ListMedicalConditions() ([]*models.MedicalCondition, error) {
	return medicalConditionService.medicalConditionDao.ListMedicalConditions()
}

func (medicalConditionService *MedicalConditionService) GetMedicalCondition(id int64) (*models.MedicalCondition, error) {
	return medicalConditionService.medicalConditionDao.GetMedicalCondition(id)
}
