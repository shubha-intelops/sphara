package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AmbulanceController struct {
	ambulanceService *services.AmbulanceService
}

func NewAmbulanceController() (*AmbulanceController, error) {
	ambulanceService, err := services.NewAmbulanceService()
	if err != nil {
		return nil, err
	}
	return &AmbulanceController{
		ambulanceService: ambulanceService,
	}, nil
}

func (ambulanceController *AmbulanceController) CreateAmbulance(context *gin.Context) {
	// validate input
	var input models.Ambulance
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger ambulance creation
	if _, err := ambulanceController.ambulanceService.CreateAmbulance(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Ambulance created successfully"})
}

func (ambulanceController *AmbulanceController) UpdateAmbulance(context *gin.Context) {
	// validate input
	var input models.Ambulance
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

	// trigger ambulance update
	if _, err := ambulanceController.ambulanceService.UpdateAmbulance(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Ambulance updated successfully"})
}

func (ambulanceController *AmbulanceController) FetchAmbulance(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger ambulance fetching
	ambulance, err := ambulanceController.ambulanceService.GetAmbulance(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, ambulance)
}

func (ambulanceController *AmbulanceController) DeleteAmbulance(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger ambulance deletion
	if err := ambulanceController.ambulanceService.DeleteAmbulance(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Ambulance deleted successfully",
	})
}

func (ambulanceController *AmbulanceController) ListAmbulances(context *gin.Context) {
	// trigger all ambulances fetching
	ambulances, err := ambulanceController.ambulanceService.ListAmbulances()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, ambulances)
}

func (*AmbulanceController) PatchAmbulance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*AmbulanceController) OptionsAmbulance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*AmbulanceController) HeadAmbulance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
