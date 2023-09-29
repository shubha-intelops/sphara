package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RobberyDao struct {
	db *gorm.DB
}

func NewRobberyDao() (*RobberyDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Robbery{})
	if err != nil {
		return nil, err
	}
	return &RobberyDao{
		db: sqlClient.DB,
	}, nil
}

func (robberyDao *RobberyDao) CreateRobbery(m *models.Robbery) (*models.Robbery, error) {
	if err := robberyDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create robbery: %v", err)
		return nil, err
	}

	log.Debugf("robbery created")
	return m, nil
}

func (robberyDao *RobberyDao) UpdateRobbery(id int64, m *models.Robbery) (*models.Robbery, error) {
	if id == 0 {
		return nil, errors.New("invalid robbery ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var robbery *models.Robbery
	if err := robberyDao.db.Where("id = ?", id).First(&robbery).Error; err != nil {
		log.Debugf("failed to find robbery for update: %v", err)
		return nil, err
	}

	if err := robberyDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update robbery: %v", err)
		return nil, err
	}
	log.Debugf("robbery updated")
	return m, nil
}

func (robberyDao *RobberyDao) DeleteRobbery(id int64) error {
	var m *models.Robbery
	if err := robberyDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete robbery: %v", err)
		return err
	}

	log.Debugf("robbery deleted")
	return nil
}

func (robberyDao *RobberyDao) ListRobberies() ([]*models.Robbery, error) {
	var robberies []*models.Robbery
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := robberyDao.db.Find(&robberies).Error; err != nil {
		log.Debugf("failed to list robberies: %v", err)
		return nil, err
	}

	log.Debugf("robbery listed")
	return robberies, nil
}

func (robberyDao *RobberyDao) GetRobbery(id int64) (*models.Robbery, error) {
	var m *models.Robbery
	if err := robberyDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get robbery: %v", err)
		return nil, err
	}
	log.Debugf("robbery retrieved")
	return m, nil
}
