package handlers

import (
	"fmt"
	"net/http"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
)

type coinHandler struct {
	coinService services.CoinService
	// authService auth.Service
}

func NewCoinHandler(coinService services.CoinService) *coinHandler {
	return &coinHandler{coinService}
}

func (h *coinHandler) GetCoins(c *gin.Context) {
	coins, err := h.coinService.FindCoins()

	if err != nil {
		response := helpers.APIResponse("Coin Failed to get", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of Coin", http.StatusOK, "success", helpers.FormatCoins(coins))
	c.JSON(http.StatusOK, response)
	return
}

func (h *coinHandler) GetCoin(c *gin.Context) {
	var input requests.GetCoinDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helpers.APIResponse("Failed to get detail coin 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	coinDetail, err := h.coinService.GetCoinBySymbol(input)

	if err != nil {
		response := helpers.APIResponse("Failed to get detail coin 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatDetailCoin(coinDetail)
	response := helpers.APIResponse("Coin Detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *coinHandler) SaveCoin(c *gin.Context) {
	var input requests.CoinInput

	currentUser := c.MustGet("currentUser").(models.User)

	userRole := currentUser.Role

	if userRole != "admin" {
		response := helpers.APIResponse("You're not administrator can not add coin", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponse("Failed to create coin", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		errorMessage := gin.H{"errors": err}
		response := helpers.APIResponse("Failed to create coin", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	input.User = currentUser
	path := fmt.Sprintf("public/logos/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		errorMessage := gin.H{"errors": err}
		response := helpers.APIResponse("Error while saving logo coin", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	coin, err := h.coinService.SaveCoin(input, path)
	if err != nil {
		response := helpers.APIResponse("Error while saving data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatCoin(coin)
	response := helpers.APIResponse("Coin success to create", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
