package medicine

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// @Summary Get all medicines
// @Tags medicines
// @Security BearerAuth
// @Produce json
// @Router /medicines [get]
func GetAll(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Create medicines
// @Tags medicines
// @Security BearerAuth
// @Produce json
// @Router /medicines [post]
func Create(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "Not implemented"}) }

// @Summary Get medicines by ID
// @Tags medicines
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medicines/{id} [get]
func GetByID(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Update medicines
// @Tags medicines
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medicines/{id} [put]
func Update(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Delete medicines
// @Tags medicines
// @Security BearerAuth
// @Param id path int true "ID"
// @Router /medicines/{id} [delete]
func Delete(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }
