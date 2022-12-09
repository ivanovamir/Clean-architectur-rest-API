package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"net/http"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var userDTO *dto.RegisterUser

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong json input",
		})
	} else {
		token, err := h.service.RegisterUser(userDTO)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}
