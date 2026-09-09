package prescription

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// @Summary Get all prescriptions
// @Tags prescriptions
// @Security BearerAuth
// @Produce json
// @Router /prescriptions [get]
func GetAll(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Create prescriptions
// @Tags prescriptions
// @Security BearerAuth
// @Produce json
// @Router /prescriptions [post]
func Create(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "Not implemented"}) }

// @Summary Get prescriptions by ID
// @Tags prescriptions
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /prescriptions/{id} [get]
func GetByID(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Update prescriptions
// @Tags prescriptions
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /prescriptions/{id} [put]
func Update(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Delete prescriptions
// @Tags prescriptions
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /prescriptions/{id} [delete]
func Delete(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }
