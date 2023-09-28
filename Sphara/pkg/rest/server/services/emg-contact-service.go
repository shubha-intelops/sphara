package services

import (
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
)

type EmgContactService struct {
	emgContactDao *daos.EmgContactDao
}

func NewEmgContactService() (*EmgContactService, error) {
	emgContactDao, err := daos.NewEmgContactDao()
	if err != nil {
		return nil, err
	}
	return &EmgContactService{
		emgContactDao: emgContactDao,
	}, nil
}

func (emgContactService *EmgContactService) CreateEmgContact(emgContact *models.EmgContact) (*models.EmgContact, error) {
	return emgContactService.emgContactDao.CreateEmgContact(emgContact)
}

func (emgContactService *EmgContactService) UpdateEmgContact(id int64, emgContact *models.EmgContact) (*models.EmgContact, error) {
	return emgContactService.emgContactDao.UpdateEmgContact(id, emgContact)
}

func (emgContactService *EmgContactService) DeleteEmgContact(id int64) error {
	return emgContactService.emgContactDao.DeleteEmgContact(id)
}

func (emgContactService *EmgContactService) ListEmgContacts() ([]*models.EmgContact, error) {
	return emgContactService.emgContactDao.ListEmgContacts()
}

func (emgContactService *EmgContactService) GetEmgContact(id int64) (*models.EmgContact, error) {
	return emgContactService.emgContactDao.GetEmgContact(id)
}
