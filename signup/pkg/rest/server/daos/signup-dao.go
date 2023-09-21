package daos

import (
	"errors"

	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SignupDao struct {
	db *gorm.DB
}

func NewSignupDao() (*SignupDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Signup{})
	if err != nil {
		return nil, err
	}
	return &SignupDao{
		db: sqlClient.DB,
	}, nil
}

func (signupDao *SignupDao) CreateSignup(m *models.Signup) (*models.Signup, error) {
	if err := signupDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create signup: %v", err)
		return nil, err
	}

	log.Debugf("signup created")
	return m, nil
}

func (signupDao *SignupDao) UpdateSignup(id int64, m *models.Signup) (*models.Signup, error) {
	if id == 0 {
		return nil, errors.New("invalid signup ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var signup *models.Signup
	if err := signupDao.db.Where("id = ?", id).First(&signup).Error; err != nil {
		log.Debugf("failed to find signup for update: %v", err)
		return nil, err
	}

	if err := signupDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update signup: %v", err)
		return nil, err
	}
	log.Debugf("signup updated")
	return m, nil
}

func (signupDao *SignupDao) DeleteSignup(id int64) error {
	var m *models.Signup
	if err := signupDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete signup: %v", err)
		return err
	}

	log.Debugf("signup deleted")
	return nil
}

func (signupDao *SignupDao) ListSignups() ([]*models.Signup, error) {
	var signups []*models.Signup
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := signupDao.db.Find(&signups).Error; err != nil {
		log.Debugf("failed to list signups: %v", err)
		return nil, err
	}

	log.Debugf("signup listed")
	return signups, nil
}

func (signupDao *SignupDao) GetSignup(id int64) (*models.Signup, error) {
	var m *models.Signup
	if err := signupDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get signup: %v", err)
		return nil, err
	}
	log.Debugf("signup retrieved")
	return m, nil
}
