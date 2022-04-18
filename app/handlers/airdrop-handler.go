package handlers

import (
	"net/http"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
)

type airdropHandler struct {
	airdropService services.AirdropService
}

func NewAirdropHandler(airdropService services.AirdropService) *airdropHandler {
	return &airdropHandler{airdropService}
}

func (h *airdropHandler) SaveAirdrop(c *gin.Context) {
	var input requests.AirdropInput

	currentUser := c.MustGet("currentUser").(models.User)

	userRole := currentUser.Role

	if userRole != "admin" {
		response := helpers.APIResponse("You're not administrator can not add coin", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.ShouldBindJSON(&input)
	newAirdrop, err := h.airdropService.CreateAirdrop(input)

	if err != nil {
		response := helpers.APIResponse("Failed to Create airdrop", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatAidrop(newAirdrop)
	response := helpers.APIResponse("Success to create airdrop", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return

}

func (h *airdropHandler) GetAirdrops(c *gin.Context) {
	airdrops, err := h.airdropService.FindAirdrop()

	if err != nil {
		response := helpers.APIResponse("Airdrops failed to get", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatAirdrops(airdrops)
	response := helpers.APIResponse("list of airdrops", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *airdropHandler) GetAirdrop(c *gin.Context) {
	var input requests.GetAirdropDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponse("Failed to get detail airdrop 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	airdropDetail, err := h.airdropService.FindAirdropByID(input)

	if err != nil {
		response := helpers.APIResponse("Failed to get detail airdrop 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatDetailAirdrop(airdropDetail)
	response := helpers.APIResponse("Airdrop detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *airdropHandler) UpdateAirdrop(c *gin.Context) {
	var inputID requests.GetAirdropDetail
	var inputStatus requests.AirdropInput

	currentUser := c.MustGet("currentUser").(models.User)
	userRole := currentUser.Role

	if userRole != "admin" {
		response := helpers.APIResponse("You're not administrator can not add coin", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.ShouldBindUri(&inputID)
	c.ShouldBindJSON(&inputStatus)

	UpdateAirdrop, err := h.airdropService.UpdateStatus(inputID, inputStatus)

	if err != nil {
		response := helpers.APIResponse("Failed to update airdrop", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatAidrop(UpdateAirdrop)
	response := helpers.APIResponse("Success to update airdrop", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
