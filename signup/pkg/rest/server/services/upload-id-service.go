package services

import (
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
)

type UploadIdService struct {
	uploadIdDao *daos.UploadIdDao
}

func NewUploadIdService() (*UploadIdService, error) {
	uploadIdDao, err := daos.NewUploadIdDao()
	if err != nil {
		return nil, err
	}
	return &UploadIdService{
		uploadIdDao: uploadIdDao,
	}, nil
}

func (uploadIdService *UploadIdService) CreateUploadId(uploadId *models.UploadId) (*models.UploadId, error) {
	return uploadIdService.uploadIdDao.CreateUploadId(uploadId)
}

func (uploadIdService *UploadIdService) UpdateUploadId(id int64, uploadId *models.UploadId) (*models.UploadId, error) {
	return uploadIdService.uploadIdDao.UpdateUploadId(id, uploadId)
}

func (uploadIdService *UploadIdService) DeleteUploadId(id int64) error {
	return uploadIdService.uploadIdDao.DeleteUploadId(id)
}

func (uploadIdService *UploadIdService) ListUploadIds() ([]*models.UploadId, error) {
	return uploadIdService.uploadIdDao.ListUploadIds()
}

func (uploadIdService *UploadIdService) GetUploadId(id int64) (*models.UploadId, error) {
	return uploadIdService.uploadIdDao.GetUploadId(id)
}
