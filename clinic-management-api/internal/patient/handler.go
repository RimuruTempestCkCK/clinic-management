package patient

import (
	"clinic-management-api/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all patients
// @Description Get a list of all patients
// @Tags patients
// @Security BearerAuth
// @Produce json
// @Success 200 {array} database.Patient
// @Router /patients [get]
func GetAll(c *gin.Context) {
	var patients []database.Patient
	database.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

// @Summary Create patient
// @Description Create a new patient
// @Tags patients
// @Security BearerAuth
// @Produce json
// @Success 201 {object} database.Patient
// @Router /patients [post]
func Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Not implemented yet"})
}

// @Summary Get patient by ID
// @Tags patients
// @Security BearerAuth
// @Produce json
// @Param id path int true "Patient ID"
// @Router /patients/{id} [get]
func GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}

// @Summary Update patient
// @Tags patients
// @Security BearerAuth
// @Produce json
// @Param id path int true "Patient ID"
// @Router /patients/{id} [put]
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}

// @Summary Delete patient
// @Tags patients
// @Security BearerAuth
// @Produce json
// @Param id path int true "Patient ID"
// @Router /patients/{id} [delete]
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}
