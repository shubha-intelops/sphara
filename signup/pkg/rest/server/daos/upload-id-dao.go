package daos

import (
	"errors"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UploadIdDao struct {
	db *gorm.DB
}

func NewUploadIdDao() (*UploadIdDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.UploadId{})
	if err != nil {
		return nil, err
	}
	return &UploadIdDao{
		db: sqlClient.DB,
	}, nil
}

func (uploadIdDao *UploadIdDao) CreateUploadId(m *models.UploadId) (*models.UploadId, error) {
	if err := uploadIdDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create uploadId: %v", err)
		return nil, err
	}

	log.Debugf("uploadId created")
	return m, nil
}

func (uploadIdDao *UploadIdDao) UpdateUploadId(id int64, m *models.UploadId) (*models.UploadId, error) {
	if id == 0 {
		return nil, errors.New("invalid uploadId ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var uploadId *models.UploadId
	if err := uploadIdDao.db.Where("id = ?", id).First(&uploadId).Error; err != nil {
		log.Debugf("failed to find uploadId for update: %v", err)
		return nil, err
	}

	if err := uploadIdDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update uploadId: %v", err)
		return nil, err
	}
	log.Debugf("uploadId updated")
	return m, nil
}

func (uploadIdDao *UploadIdDao) DeleteUploadId(id int64) error {
	var m *models.UploadId
	if err := uploadIdDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete uploadId: %v", err)
		return err
	}

	log.Debugf("uploadId deleted")
	return nil
}

func (uploadIdDao *UploadIdDao) ListUploadIds() ([]*models.UploadId, error) {
	var uploadIds []*models.UploadId
	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := uploadIdDao.db.Find(&uploadIds).Error; err != nil {
		log.Debugf("failed to list uploadIds: %v", err)
		return nil, err
	}

	log.Debugf("uploadId listed")
	return uploadIds, nil
}

func (uploadIdDao *UploadIdDao) GetUploadId(id int64) (*models.UploadId, error) {
	var m *models.UploadId
	if err := uploadIdDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get uploadId: %v", err)
		return nil, err
	}
	log.Debugf("uploadId retrieved")
	return m, nil
}
