package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PanicController struct {
	panicService *services.PanicService
}

func NewPanicController() (*PanicController, error) {
	panicService, err := services.NewPanicService()
	if err != nil {
		return nil, err
	}
	return &PanicController{
		panicService: panicService,
	}, nil
}

func (panicController *PanicController) CreatePanic(context *gin.Context) {
	// validate input
	var input models.Panic
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger panic creation
	if _, err := panicController.panicService.CreatePanic(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Panic created successfully"})
}

func (panicController *PanicController) UpdatePanic(context *gin.Context) {
	// validate input
	var input models.Panic
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

	// trigger panic update
	if _, err := panicController.panicService.UpdatePanic(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Panic updated successfully"})
}

func (panicController *PanicController) FetchPanic(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger panic fetching
	panic, err := panicController.panicService.GetPanic(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, panic)
}

func (panicController *PanicController) DeletePanic(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger panic deletion
	if err := panicController.panicService.DeletePanic(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Panic deleted successfully",
	})
}

func (panicController *PanicController) ListPanics(context *gin.Context) {
	// trigger all panics fetching
	panics, err := panicController.panicService.ListPanics()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, panics)
}

func (*PanicController) PatchPanic(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*PanicController) OptionsPanic(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*PanicController) HeadPanic(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
