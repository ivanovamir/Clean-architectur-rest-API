package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"net/http"
)

func (h *Handler) GetAllCategories(c *gin.Context) {

	allCategoriesDTO, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": []string{}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allCategoriesDTO,
		})
	}
}

func (h *Handler) GetCategoriesById(c *gin.Context) {
	categoryId := c.Query("cat_id")
	categoryIdDTO := &dto.CategoryParams{
		ID: categoryId,
	}
	categoryDTO, err := h.service.GetCategoriesById(categoryIdDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": []string{}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": categoryDTO,
		})
	}
}
