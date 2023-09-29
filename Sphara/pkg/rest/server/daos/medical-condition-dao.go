package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MedicalConditionDao struct {
	db *gorm.DB
}

func NewMedicalConditionDao() (*MedicalConditionDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.MedicalCondition{})
	if err != nil {
		return nil, err
	}
	return &MedicalConditionDao{
		db: sqlClient.DB,
	}, nil
}

func (medicalConditionDao *MedicalConditionDao) CreateMedicalCondition(m *models.MedicalCondition) (*models.MedicalCondition, error) {
	if err := medicalConditionDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create medicalCondition: %v", err)
		return nil, err
	}

	log.Debugf("medicalCondition created")
	return m, nil
}

func (medicalConditionDao *MedicalConditionDao) UpdateMedicalCondition(id int64, m *models.MedicalCondition) (*models.MedicalCondition, error) {
	if id == 0 {
		return nil, errors.New("invalid medicalCondition ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var medicalCondition *models.MedicalCondition
	if err := medicalConditionDao.db.Where("id = ?", id).First(&medicalCondition).Error; err != nil {
		log.Debugf("failed to find medicalCondition for update: %v", err)
		return nil, err
	}

	if err := medicalConditionDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update medicalCondition: %v", err)
		return nil, err
	}
	log.Debugf("medicalCondition updated")
	return m, nil
}

func (medicalConditionDao *MedicalConditionDao) DeleteMedicalCondition(id int64) error {
	var m *models.MedicalCondition
	if err := medicalConditionDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete medicalCondition: %v", err)
		return err
	}

	log.Debugf("medicalCondition deleted")
	return nil
}

func (medicalConditionDao *MedicalConditionDao) ListMedicalConditions() ([]*models.MedicalCondition, error) {
	var medicalConditions []*models.MedicalCondition
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := medicalConditionDao.db.Find(&medicalConditions).Error; err != nil {
		log.Debugf("failed to list medicalConditions: %v", err)
		return nil, err
	}

	log.Debugf("medicalCondition listed")
	return medicalConditions, nil
}

func (medicalConditionDao *MedicalConditionDao) GetMedicalCondition(id int64) (*models.MedicalCondition, error) {
	var m *models.MedicalCondition
	if err := medicalConditionDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get medicalCondition: %v", err)
		return nil, err
	}
	log.Debugf("medicalCondition retrieved")
	return m, nil
}
