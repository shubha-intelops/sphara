package services

import (
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
)

type MedicalInsuranceService struct {
	medicalInsuranceDao *daos.MedicalInsuranceDao
}

func NewMedicalInsuranceService() (*MedicalInsuranceService, error) {
	medicalInsuranceDao, err := daos.NewMedicalInsuranceDao()
	if err != nil {
		return nil, err
	}
	return &MedicalInsuranceService{
		medicalInsuranceDao: medicalInsuranceDao,
	}, nil
}

func (medicalInsuranceService *MedicalInsuranceService) CreateMedicalInsurance(medicalInsurance *models.MedicalInsurance) (*models.MedicalInsurance, error) {
	return medicalInsuranceService.medicalInsuranceDao.CreateMedicalInsurance(medicalInsurance)
}

func (medicalInsuranceService *MedicalInsuranceService) UpdateMedicalInsurance(id int64, medicalInsurance *models.MedicalInsurance) (*models.MedicalInsurance, error) {
	return medicalInsuranceService.medicalInsuranceDao.UpdateMedicalInsurance(id, medicalInsurance)
}

func (medicalInsuranceService *MedicalInsuranceService) DeleteMedicalInsurance(id int64) error {
	return medicalInsuranceService.medicalInsuranceDao.DeleteMedicalInsurance(id)
}

func (medicalInsuranceService *MedicalInsuranceService) ListMedicalInsurances() ([]*models.MedicalInsurance, error) {
	return medicalInsuranceService.medicalInsuranceDao.ListMedicalInsurances()
}

func (medicalInsuranceService *MedicalInsuranceService) GetMedicalInsurance(id int64) (*models.MedicalInsurance, error) {
	return medicalInsuranceService.medicalInsuranceDao.GetMedicalInsurance(id)
}
