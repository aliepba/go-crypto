package handlers

import (
	"net/http"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
)

type metadataHandler struct {
	metadataService services.MetadataService
}

func NewMetadataHandler(metadataService services.MetadataService) *metadataHandler {
	return &metadataHandler{metadataService}
}

func (h *metadataHandler) GetMetadata(c *gin.Context) {
	metadata, err := h.metadataService.FindMetadata()

	if err != nil {
		response := helpers.APIResponse("Metadata failed to get", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatMetadataCoins(metadata)
	response := helpers.APIResponse("List of metadata", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *metadataHandler) SaveMetadata(c *gin.Context) {
	var input requests.MetadataInput

	currentUser := c.MustGet("currentUser").(models.User)

	userRole := currentUser.Role

	if userRole != "admin" {
		response := helpers.APIResponse("You're not administrator can not add coin", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.ShouldBindJSON(&input)
	input.User = currentUser

	newMetadata, err := h.metadataService.SaveMetadata(input)

	if err != nil {
		response := helpers.APIResponse("Failed to Create Metadata", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatMetadata(newMetadata)
	response := helpers.APIResponse("Success to create metadata", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
