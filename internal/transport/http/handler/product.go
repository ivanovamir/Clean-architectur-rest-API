package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"net/http"
)

func (h *Handler) GetAllProductsByParams(c *gin.Context) {

	allParams := h.ParseUrlParams(c)
	allProducts, err := h.service.GetAllProductsByParams(allParams)
	if err != nil || allProducts == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allProducts,
		})
	}
}

func (h *Handler) GetProductDetail(c *gin.Context) {
	productId := c.Query("prod_id")
	productInfo, err := h.service.GetProductDetail(productId)
	if err != nil && productInfo == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []*dto.ProductInformation{productInfo},
		})
	}
}
