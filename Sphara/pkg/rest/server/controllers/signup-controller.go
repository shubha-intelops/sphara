package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type SignupController struct {
	signupService *services.SignupService
}

func NewSignupController() (*SignupController, error) {
	signupService, err := services.NewSignupService()
	if err != nil {
		return nil, err
	}
	return &SignupController{
		signupService: signupService,
	}, nil
}

func (signupController *SignupController) CreateSignup(context *gin.Context) {
	// validate input
	var input models.Signup
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger signup creation
	if _, err := signupController.signupService.CreateSignup(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signup created successfully"})
}

func (signupController *SignupController) UpdateSignup(context *gin.Context) {
	// validate input
	var input models.Signup
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger signup update
	if _, err := signupController.signupService.UpdateSignup(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Signup updated successfully"})
}

func (signupController *SignupController) FetchSignup(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger signup fetching
	signup, err := signupController.signupService.GetSignup(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, signup)
}

func (signupController *SignupController) DeleteSignup(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger signup deletion
	if err := signupController.signupService.DeleteSignup(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Signup deleted successfully",
	})
}

func (signupController *SignupController) ListSignups(context *gin.Context) {
	// trigger all signups fetching
	signups, err := signupController.signupService.ListSignups()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, signups)
}

func (*SignupController) PatchSignup(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*SignupController) OptionsSignup(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*SignupController) HeadSignup(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
