package services

import (
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/daos"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
)

type SignupService struct {
	signupDao *daos.SignupDao
}

func NewSignupService() (*SignupService, error) {
	signupDao, err := daos.NewSignupDao()
	if err != nil {
		return nil, err
	}
	return &SignupService{
		signupDao: signupDao,
	}, nil
}

func (signupService *SignupService) CreateSignup(signup *models.Signup) (*models.Signup, error) {
	return signupService.signupDao.CreateSignup(signup)
}

func (signupService *SignupService) UpdateSignup(id int64, signup *models.Signup) (*models.Signup, error) {
	return signupService.signupDao.UpdateSignup(id, signup)
}

func (signupService *SignupService) DeleteSignup(id int64) error {
	return signupService.signupDao.DeleteSignup(id)
}

func (signupService *SignupService) ListSignups() ([]*models.Signup, error) {
	return signupService.signupDao.ListSignups()
}

func (signupService *SignupService) GetSignup(id int64) (*models.Signup, error) {
	return signupService.signupDao.GetSignup(id)
}
