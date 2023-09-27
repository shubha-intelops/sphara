package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type EmgContactController struct {
	emgContactService *services.EmgContactService
}

func NewEmgContactController() (*EmgContactController, error) {
	emgContactService, err := services.NewEmgContactService()
	if err != nil {
		return nil, err
	}
	return &EmgContactController{
		emgContactService: emgContactService,
	}, nil
}

func (emgContactController *EmgContactController) CreateEmgContact(context *gin.Context) {
	// validate input
	var input models.EmgContact
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger emgContact creation
	if _, err := emgContactController.emgContactService.CreateEmgContact(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "EmgContact created successfully"})
}

func (emgContactController *EmgContactController) UpdateEmgContact(context *gin.Context) {
	// validate input
	var input models.EmgContact
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

	// trigger emgContact update
	if _, err := emgContactController.emgContactService.UpdateEmgContact(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "EmgContact updated successfully"})
}

func (emgContactController *EmgContactController) FetchEmgContact(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger emgContact fetching
	emgContact, err := emgContactController.emgContactService.GetEmgContact(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, emgContact)
}

func (emgContactController *EmgContactController) DeleteEmgContact(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger emgContact deletion
	if err := emgContactController.emgContactService.DeleteEmgContact(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "EmgContact deleted successfully",
	})
}

func (emgContactController *EmgContactController) ListEmgContacts(context *gin.Context) {
	// trigger all emgContacts fetching
	emgContacts, err := emgContactController.emgContactService.ListEmgContacts()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, emgContacts)
}

func (*EmgContactController) PatchEmgContact(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*EmgContactController) OptionsEmgContact(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*EmgContactController) HeadEmgContact(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
