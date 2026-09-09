package auth

import (
	"clinic-management-api/internal/database"
	"clinic-management-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=ADMIN DOCTOR STAFF PATIENT"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// @Summary Register User
// @Description Register a new user (ADMIN, DOCTOR, STAFF, PATIENT)
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterInput true "Registration Data"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := database.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": user.ID})
}

// @Summary Login User
// @Description Login and get JWT Token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginInput true "Login Data"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user database.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Refresh Token
// @Tags auth
// @Produce json
// @Router /auth/refresh [post]
func Refresh(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }

// @Summary Logout
// @Tags auth
// @Produce json
// @Router /auth/logout [post]
func Logout(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Not implemented"}) }
