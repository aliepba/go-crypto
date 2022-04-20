package handlers

import (
	"net/http"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
)

type fiatHandler struct {
	fiatService services.FiatService
}

func NewFiatHandler(fiatService services.FiatService) *fiatHandler {
	return &fiatHandler{fiatService}
}

func (h *fiatHandler) SaveFiat(c *gin.Context) {
	var input requests.FiatInput

	currentUser := c.MustGet("currentUser").(models.User)
	userRole := currentUser.Role

	if userRole != "admin" {
		response := helpers.APIResponse("You're not administrator can not add coin", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponse("Failed to create fiat", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newFiat, err := h.fiatService.SaveFiat(input)

	if err != nil {
		response := helpers.APIResponse("Failed to Create fiat", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatFiat(newFiat)
	response := helpers.APIResponse("Success to create fiat", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
