package handlers

import (
	"net/http"

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
