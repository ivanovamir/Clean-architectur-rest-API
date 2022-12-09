package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	//TODO rework it

	api := router.Group("/api")
	{
		auth := api.Group("/auth", h.UserIdentity)
		{
			auth.POST("/sign-up", h.RegisterUser)
			//auth.POST("/sign-in", h.AuthenticateUser)
		}
		categories := api.Group("/category")
		{
			categories.GET("/", h.GetAllCategories)
		}
		api.GET("/category_detail", h.GetCategoriesById)
		products := api.Group("/products")
		{
			products.GET("/", h.GetAllProductsByParams)

		}
		api.GET("/product_detail", h.GetProductDetail)
	}
	return router
}
