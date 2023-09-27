package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type MedicalConditionController struct {
	medicalConditionService *services.MedicalConditionService
}

func NewMedicalConditionController() (*MedicalConditionController, error) {
	medicalConditionService, err := services.NewMedicalConditionService()
	if err != nil {
		return nil, err
	}
	return &MedicalConditionController{
		medicalConditionService: medicalConditionService,
	}, nil
}

func (medicalConditionController *MedicalConditionController) CreateMedicalCondition(context *gin.Context) {
	// validate input
	var input models.MedicalCondition
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalCondition creation
	if _, err := medicalConditionController.medicalConditionService.CreateMedicalCondition(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "MedicalCondition created successfully"})
}

func (medicalConditionController *MedicalConditionController) UpdateMedicalCondition(context *gin.Context) {
	// validate input
	var input models.MedicalCondition
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

	// trigger medicalCondition update
	if _, err := medicalConditionController.medicalConditionService.UpdateMedicalCondition(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "MedicalCondition updated successfully"})
}

func (medicalConditionController *MedicalConditionController) FetchMedicalCondition(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalCondition fetching
	medicalCondition, err := medicalConditionController.medicalConditionService.GetMedicalCondition(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, medicalCondition)
}

func (medicalConditionController *MedicalConditionController) DeleteMedicalCondition(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger medicalCondition deletion
	if err := medicalConditionController.medicalConditionService.DeleteMedicalCondition(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "MedicalCondition deleted successfully",
	})
}

func (medicalConditionController *MedicalConditionController) ListMedicalConditions(context *gin.Context) {
	// trigger all medicalConditions fetching
	medicalConditions, err := medicalConditionController.medicalConditionService.ListMedicalConditions()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, medicalConditions)
}

func (*MedicalConditionController) PatchMedicalCondition(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*MedicalConditionController) OptionsMedicalCondition(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*MedicalConditionController) HeadMedicalCondition(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
