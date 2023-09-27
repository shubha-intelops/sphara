package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type MedicalInsuranceController struct {
	medicalInsuranceService *services.MedicalInsuranceService
}

func NewMedicalInsuranceController() (*MedicalInsuranceController, error) {
	medicalInsuranceService, err := services.NewMedicalInsuranceService()
	if err != nil {
		return nil, err
	}
	return &MedicalInsuranceController{
		medicalInsuranceService: medicalInsuranceService,
	}, nil
}

func (medicalInsuranceController *MedicalInsuranceController) CreateMedicalInsurance(context *gin.Context) {
	// validate input
	var input models.MedicalInsurance
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalInsurance creation
	if _, err := medicalInsuranceController.medicalInsuranceService.CreateMedicalInsurance(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "MedicalInsurance created successfully"})
}

func (medicalInsuranceController *MedicalInsuranceController) UpdateMedicalInsurance(context *gin.Context) {
	// validate input
	var input models.MedicalInsurance
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

	// trigger medicalInsurance update
	if _, err := medicalInsuranceController.medicalInsuranceService.UpdateMedicalInsurance(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "MedicalInsurance updated successfully"})
}

func (medicalInsuranceController *MedicalInsuranceController) FetchMedicalInsurance(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalInsurance fetching
	medicalInsurance, err := medicalInsuranceController.medicalInsuranceService.GetMedicalInsurance(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, medicalInsurance)
}

func (medicalInsuranceController *MedicalInsuranceController) DeleteMedicalInsurance(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalInsurance deletion
	if err := medicalInsuranceController.medicalInsuranceService.DeleteMedicalInsurance(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "MedicalInsurance deleted successfully",
	})
}

func (medicalInsuranceController *MedicalInsuranceController) ListMedicalInsurances(context *gin.Context) {
	// trigger all medicalInsurances fetching
	medicalInsurances, err := medicalInsuranceController.medicalInsuranceService.ListMedicalInsurances()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, medicalInsurances)
}

func (*MedicalInsuranceController) PatchMedicalInsurance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*MedicalInsuranceController) OptionsMedicalInsurance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*MedicalInsuranceController) HeadMedicalInsurance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
