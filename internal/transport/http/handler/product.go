package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetAllProductsByParams(c *gin.Context) {

	allParams := h.ParseUrlParams(c)
	allProducts, err := h.service.GetAllProductsByParams(allParams)
	if err != nil {
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []*dto.ProductInformation{productInfo},
		})
	}
}

func (h *Handler) ParseUrlParams(c *gin.Context) *dto.Params {
	allCategoriesId := h.service.ParseUrlParams()

	limit := 0
	page := 1
	search := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "cat_id":
			allCategoriesId = toIntArray(queryValue)
		case "s":
			search = queryValue
		}
	}
	return &dto.Params{
		Limit:  limit,
		Page:   page,
		CatId:  allCategoriesId,
		Search: search,
	}
}

func toIntArray(str string) []int {
	chunks := strings.Split(str, ",")

	var res []int
	for _, c := range chunks {
		i, _ := strconv.Atoi(c) // error handling omitted for concision
		res = append(res, i)
	}

	return res
}
