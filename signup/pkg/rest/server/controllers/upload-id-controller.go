package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/models"
	"github.com/shubha-intelops/sphara/signup/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
)

type UploadIdController struct {
	uploadIdService *services.UploadIdService
}

func NewUploadIdController() (*UploadIdController, error) {
	uploadIdService, err := services.NewUploadIdService()
	if err != nil {
		return nil, err
	}
	return &UploadIdController{
		uploadIdService: uploadIdService,
	}, nil
}

func (uploadIdController *UploadIdController) CreateUploadId(context *gin.Context) {
	var input models.UploadId

	input.Id_number = context.PostForm("idNumber")
	input.Id_number = context.PostForm("idType")
	inputFile, err := context.FormFile("file")
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dst := "./" + inputFile.Filename
	if err := context.SaveUploadedFile(inputFile, dst); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.File = dst
	// trigger uploadId creation
	if _, err := uploadIdController.uploadIdService.CreateUploadId(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "UploadId created successfully", "file": dst})
}

func (uploadIdController *UploadIdController) UpdateUploadId(context *gin.Context) {
	// validate input
	var input models.UploadId
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

	// trigger uploadId update
	if _, err := uploadIdController.uploadIdService.UpdateUploadId(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "UploadId updated successfully"})
}

func (uploadIdController *UploadIdController) FetchUploadId(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger uploadId fetching
	uploadId, err := uploadIdController.uploadIdService.GetUploadId(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, uploadId)
}

func (uploadIdController *UploadIdController) DeleteUploadId(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger uploadId deletion
	if err := uploadIdController.uploadIdService.DeleteUploadId(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "UploadId deleted successfully",
	})
}

func (uploadIdController *UploadIdController) ListUploadIds(context *gin.Context) {
	// trigger all uploadIds fetching
	uploadIds, err := uploadIdController.uploadIdService.ListUploadIds()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, uploadIds)
}

func (*UploadIdController) PatchUploadId(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*UploadIdController) OptionsUploadId(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*UploadIdController) HeadUploadId(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
