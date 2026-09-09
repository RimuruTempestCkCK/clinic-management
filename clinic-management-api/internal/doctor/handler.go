package doctor

import (
	"clinic-management-api/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all doctors
// @Tags doctors
// @Security BearerAuth
// @Produce json
// @Success 200 {array} database.Doctor
// @Router /doctors [get]
func GetAll(c *gin.Context) {
	var doctors []database.Doctor
	database.DB.Find(&doctors)
	c.JSON(http.StatusOK, doctors)
}

// @Summary Create doctor
// @Tags doctors
// @Security BearerAuth
// @Produce json
// @Router /doctors [post]
func Create(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "Not implemented"}) }

// @Summary Get doctor by ID
// @Tags doctors
// @Security BearerAuth
// @Param id path int true "Doctor ID"
// @Router /doctors/{id} [get]
func GetByID(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Update doctor
// @Tags doctors
// @Security BearerAuth
// @Param id path int true "Doctor ID"
// @Router /doctors/{id} [put]
func Update(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Delete doctor
// @Tags doctors
// @Security BearerAuth
// @Param id path int true "Doctor ID"
// @Router /doctors/{id} [delete]
func Delete(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }
