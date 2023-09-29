package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PanicDao struct {
	db *gorm.DB
}

func NewPanicDao() (*PanicDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Panic{})
	if err != nil {
		return nil, err
	}
	return &PanicDao{
		db: sqlClient.DB,
	}, nil
}

func (panicDao *PanicDao) CreatePanic(m *models.Panic) (*models.Panic, error) {
	if err := panicDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create panic: %v", err)
		return nil, err
	}

	log.Debugf("panic created")
	return m, nil
}

func (panicDao *PanicDao) UpdatePanic(id int64, m *models.Panic) (*models.Panic, error) {
	if id == 0 {
		return nil, errors.New("invalid panic ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var panic *models.Panic
	if err := panicDao.db.Where("id = ?", id).First(&panic).Error; err != nil {
		log.Debugf("failed to find panic for update: %v", err)
		return nil, err
	}

	if err := panicDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update panic: %v", err)
		return nil, err
	}
	log.Debugf("panic updated")
	return m, nil
}

func (panicDao *PanicDao) DeletePanic(id int64) error {
	var m *models.Panic
	if err := panicDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete panic: %v", err)
		return err
	}

	log.Debugf("panic deleted")
	return nil
}

func (panicDao *PanicDao) ListPanics() ([]*models.Panic, error) {
	var panics []*models.Panic
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := panicDao.db.Find(&panics).Error; err != nil {
		log.Debugf("failed to list panics: %v", err)
		return nil, err
	}

	log.Debugf("panic listed")
	return panics, nil
}

func (panicDao *PanicDao) GetPanic(id int64) (*models.Panic, error) {
	var m *models.Panic
	if err := panicDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get panic: %v", err)
		return nil, err
	}
	log.Debugf("panic retrieved")
	return m, nil
}
