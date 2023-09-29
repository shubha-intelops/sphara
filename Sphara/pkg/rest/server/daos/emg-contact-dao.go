package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EmgContactDao struct {
	db *gorm.DB
}

func NewEmgContactDao() (*EmgContactDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.EmgContact{})
	if err != nil {
		return nil, err
	}
	return &EmgContactDao{
		db: sqlClient.DB,
	}, nil
}

func (emgContactDao *EmgContactDao) CreateEmgContact(m *models.EmgContact) (*models.EmgContact, error) {
	if err := emgContactDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create emgContact: %v", err)
		return nil, err
	}

	log.Debugf("emgContact created")
	return m, nil
}

func (emgContactDao *EmgContactDao) UpdateEmgContact(id int64, m *models.EmgContact) (*models.EmgContact, error) {
	if id == 0 {
		return nil, errors.New("invalid emgContact ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var emgContact *models.EmgContact
	if err := emgContactDao.db.Where("id = ?", id).First(&emgContact).Error; err != nil {
		log.Debugf("failed to find emgContact for update: %v", err)
		return nil, err
	}

	if err := emgContactDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update emgContact: %v", err)
		return nil, err
	}
	log.Debugf("emgContact updated")
	return m, nil
}

func (emgContactDao *EmgContactDao) DeleteEmgContact(id int64) error {
	var m *models.EmgContact
	if err := emgContactDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete emgContact: %v", err)
		return err
	}

	log.Debugf("emgContact deleted")
	return nil
}

func (emgContactDao *EmgContactDao) ListEmgContacts() ([]*models.EmgContact, error) {
	var emgContacts []*models.EmgContact
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := emgContactDao.db.Find(&emgContacts).Error; err != nil {
		log.Debugf("failed to list emgContacts: %v", err)
		return nil, err
	}

	log.Debugf("emgContact listed")
	return emgContacts, nil
}

func (emgContactDao *EmgContactDao) GetEmgContact(id int64) (*models.EmgContact, error) {
	var m *models.EmgContact
	if err := emgContactDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get emgContact: %v", err)
		return nil, err
	}
	log.Debugf("emgContact retrieved")
	return m, nil
}
