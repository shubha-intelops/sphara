package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FireDao struct {
	db *gorm.DB
}

func NewFireDao() (*FireDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Fire{})
	if err != nil {
		return nil, err
	}
	return &FireDao{
		db: sqlClient.DB,
	}, nil
}

func (fireDao *FireDao) CreateFire(m *models.Fire) (*models.Fire, error) {
	if err := fireDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create fire: %v", err)
		return nil, err
	}

	log.Debugf("fire created")
	return m, nil
}

func (fireDao *FireDao) UpdateFire(id int64, m *models.Fire) (*models.Fire, error) {
	if id == 0 {
		return nil, errors.New("invalid fire ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var fire *models.Fire
	if err := fireDao.db.Where("id = ?", id).First(&fire).Error; err != nil {
		log.Debugf("failed to find fire for update: %v", err)
		return nil, err
	}

	if err := fireDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update fire: %v", err)
		return nil, err
	}
	log.Debugf("fire updated")
	return m, nil
}

func (fireDao *FireDao) DeleteFire(id int64) error {
	var m *models.Fire
	if err := fireDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete fire: %v", err)
		return err
	}

	log.Debugf("fire deleted")
	return nil
}

func (fireDao *FireDao) ListFires() ([]*models.Fire, error) {
	var fires []*models.Fire
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := fireDao.db.Find(&fires).Error; err != nil {
		log.Debugf("failed to list fires: %v", err)
		return nil, err
	}

	log.Debugf("fire listed")
	return fires, nil
}

func (fireDao *FireDao) GetFire(id int64) (*models.Fire, error) {
	var m *models.Fire
	if err := fireDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get fire: %v", err)
		return nil, err
	}
	log.Debugf("fire retrieved")
	return m, nil
}
