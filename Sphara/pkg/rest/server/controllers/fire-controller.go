package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type FireController struct {
	fireService *services.FireService
}

func NewFireController() (*FireController, error) {
	fireService, err := services.NewFireService()
	if err != nil {
		return nil, err
	}
	return &FireController{
		fireService: fireService,
	}, nil
}

func (fireController *FireController) CreateFire(context *gin.Context) {
	// validate input
	var input models.Fire
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger fire creation
	if _, err := fireController.fireService.CreateFire(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Fire created successfully"})
}

func (fireController *FireController) UpdateFire(context *gin.Context) {
	// validate input
	var input models.Fire
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

	// trigger fire update
	if _, err := fireController.fireService.UpdateFire(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Fire updated successfully"})
}

func (fireController *FireController) FetchFire(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger fire fetching
	fire, err := fireController.fireService.GetFire(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, fire)
}

func (fireController *FireController) DeleteFire(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger fire deletion
	if err := fireController.fireService.DeleteFire(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Fire deleted successfully",
	})
}

func (fireController *FireController) ListFires(context *gin.Context) {
	// trigger all fires fetching
	fires, err := fireController.fireService.ListFires()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, fires)
}

func (*FireController) PatchFire(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*FireController) OptionsFire(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*FireController) HeadFire(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
