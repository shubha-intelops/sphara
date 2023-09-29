package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/sphara/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RobberyController struct {
	robberyService *services.RobberyService
}

func NewRobberyController() (*RobberyController, error) {
	robberyService, err := services.NewRobberyService()
	if err != nil {
		return nil, err
	}
	return &RobberyController{
		robberyService: robberyService,
	}, nil
}

func (robberyController *RobberyController) CreateRobbery(context *gin.Context) {
	// validate input
	var input models.Robbery
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger robbery creation
	if _, err := robberyController.robberyService.CreateRobbery(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Robbery created successfully"})
}

func (robberyController *RobberyController) UpdateRobbery(context *gin.Context) {
	// validate input
	var input models.Robbery
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

	// trigger robbery update
	if _, err := robberyController.robberyService.UpdateRobbery(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Robbery updated successfully"})
}

func (robberyController *RobberyController) FetchRobbery(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger robbery fetching
	robbery, err := robberyController.robberyService.GetRobbery(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, robbery)
}

func (robberyController *RobberyController) DeleteRobbery(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger robbery deletion
	if err := robberyController.robberyService.DeleteRobbery(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Robbery deleted successfully",
	})
}

func (robberyController *RobberyController) ListRobberies(context *gin.Context) {
	// trigger all robberies fetching
	robberies, err := robberyController.robberyService.ListRobberies()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, robberies)
}

func (*RobberyController) PatchRobbery(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*RobberyController) OptionsRobbery(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*RobberyController) HeadRobbery(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
