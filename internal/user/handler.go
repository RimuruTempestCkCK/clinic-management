package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Tags users
// @Security BearerAuth
// @Produce json
// @Router /users [get]
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
}
