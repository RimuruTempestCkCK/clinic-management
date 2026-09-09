package medical_record

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all medical-records
// @Tags medical-records
// @Security BearerAuth
// @Produce json
// @Router /medical-records [get]
func GetAll(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Create medical-records
// @Tags medical-records
// @Security BearerAuth
// @Produce json
// @Router /medical-records [post]
func Create(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "Not implemented"}) }

// @Summary Get medical-records by ID
// @Tags medical-records
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medical-records/{id} [get]
func GetByID(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Update medical-records
// @Tags medical-records
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medical-records/{id} [put]
func Update(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Delete medical-records
// @Tags medical-records
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medical-records/{id} [delete]
func Delete(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }
