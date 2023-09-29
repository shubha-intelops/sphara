package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MedicalInsuranceDao struct {
	db *gorm.DB
}

func NewMedicalInsuranceDao() (*MedicalInsuranceDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.MedicalInsurance{})
	if err != nil {
		return nil, err
	}
	return &MedicalInsuranceDao{
		db: sqlClient.DB,
	}, nil
}

func (medicalInsuranceDao *MedicalInsuranceDao) CreateMedicalInsurance(m *models.MedicalInsurance) (*models.MedicalInsurance, error) {
	if err := medicalInsuranceDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create medicalInsurance: %v", err)
		return nil, err
	}

	log.Debugf("medicalInsurance created")
	return m, nil
}

func (medicalInsuranceDao *MedicalInsuranceDao) UpdateMedicalInsurance(id int64, m *models.MedicalInsurance) (*models.MedicalInsurance, error) {
	if id == 0 {
		return nil, errors.New("invalid medicalInsurance ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var medicalInsurance *models.MedicalInsurance
	if err := medicalInsuranceDao.db.Where("id = ?", id).First(&medicalInsurance).Error; err != nil {
		log.Debugf("failed to find medicalInsurance for update: %v", err)
		return nil, err
	}

	if err := medicalInsuranceDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update medicalInsurance: %v", err)
		return nil, err
	}
	log.Debugf("medicalInsurance updated")
	return m, nil
}

func (medicalInsuranceDao *MedicalInsuranceDao) DeleteMedicalInsurance(id int64) error {
	var m *models.MedicalInsurance
	if err := medicalInsuranceDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete medicalInsurance: %v", err)
		return err
	}

	log.Debugf("medicalInsurance deleted")
	return nil
}

func (medicalInsuranceDao *MedicalInsuranceDao) ListMedicalInsurances() ([]*models.MedicalInsurance, error) {
	var medicalInsurances []*models.MedicalInsurance
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := medicalInsuranceDao.db.Find(&medicalInsurances).Error; err != nil {
		log.Debugf("failed to list medicalInsurances: %v", err)
		return nil, err
	}

	log.Debugf("medicalInsurance listed")
	return medicalInsurances, nil
}

func (medicalInsuranceDao *MedicalInsuranceDao) GetMedicalInsurance(id int64) (*models.MedicalInsurance, error) {
	var m *models.MedicalInsurance
	if err := medicalInsuranceDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get medicalInsurance: %v", err)
		return nil, err
	}
	log.Debugf("medicalInsurance retrieved")
	return m, nil
}
