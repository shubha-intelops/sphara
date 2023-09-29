package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AmbulanceDao struct {
	db *gorm.DB
}

func NewAmbulanceDao() (*AmbulanceDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Ambulance{})
	if err != nil {
		return nil, err
	}
	return &AmbulanceDao{
		db: sqlClient.DB,
	}, nil
}

func (ambulanceDao *AmbulanceDao) CreateAmbulance(m *models.Ambulance) (*models.Ambulance, error) {
	if err := ambulanceDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create ambulance: %v", err)
		return nil, err
	}

	log.Debugf("ambulance created")
	return m, nil
}

func (ambulanceDao *AmbulanceDao) UpdateAmbulance(id int64, m *models.Ambulance) (*models.Ambulance, error) {
	if id == 0 {
		return nil, errors.New("invalid ambulance ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var ambulance *models.Ambulance
	if err := ambulanceDao.db.Where("id = ?", id).First(&ambulance).Error; err != nil {
		log.Debugf("failed to find ambulance for update: %v", err)
		return nil, err
	}

	if err := ambulanceDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update ambulance: %v", err)
		return nil, err
	}
	log.Debugf("ambulance updated")
	return m, nil
}

func (ambulanceDao *AmbulanceDao) DeleteAmbulance(id int64) error {
	var m *models.Ambulance
	if err := ambulanceDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete ambulance: %v", err)
		return err
	}

	log.Debugf("ambulance deleted")
	return nil
}

func (ambulanceDao *AmbulanceDao) ListAmbulances() ([]*models.Ambulance, error) {
	var ambulances []*models.Ambulance
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := ambulanceDao.db.Find(&ambulances).Error; err != nil {
		log.Debugf("failed to list ambulances: %v", err)
		return nil, err
	}

	log.Debugf("ambulance listed")
	return ambulances, nil
}

func (ambulanceDao *AmbulanceDao) GetAmbulance(id int64) (*models.Ambulance, error) {
	var m *models.Ambulance
	if err := ambulanceDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get ambulance: %v", err)
		return nil, err
	}
	log.Debugf("ambulance retrieved")
	return m, nil
}
