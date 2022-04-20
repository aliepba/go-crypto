package handlers

import (
	"net/http"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) SaveCategory(c *gin.Context) {
	var input requests.CategoryInput

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
		response := helpers.APIResponse("Failed to create category", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.categoryService.SaveCategory(input)

	if err != nil {
		response := helpers.APIResponse("Failed to Create Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatCategory(newCategory)
	response := helpers.APIResponse("Success to create category", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.FindCategory()

	if err != nil {
		response := helpers.APIResponse("Categories failed to get", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("list of Category", http.StatusOK, "success", categories)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) GetByCategory(c *gin.Context) {
	var input requests.GetCategoryInput

	c.ShouldBindUri(&input)

	category, err := h.categoryService.GetCoinByCategory(input)

	if err != nil {
		response := helpers.APIResponse("Failed to get category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helpers.FormatByCategory(category)
	response := helpers.APIResponse("Category", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
