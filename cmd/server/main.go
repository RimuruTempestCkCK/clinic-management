package main

import (
	"clinic-management-api/internal/appointment"
	"clinic-management-api/internal/auth"
	"clinic-management-api/internal/database"
	"clinic-management-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if exists
	_ = godotenv.Load()

	// Connect Database
	database.ConnectDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", auth.Register)
			authRoutes.POST("/login", auth.Login)
			// TODO: Refresh & Logout
		}

		// Example of protected routes
		patients := v1.Group("/patients")
		patients.Use(middleware.AuthMiddleware())
		{
			patients.GET("", middleware.RoleMiddleware("ADMIN", "DOCTOR", "STAFF"), func(c *gin.Context) {
				var patientList []database.Patient
				database.DB.Find(&patientList)
				c.JSON(200, patientList)
			})
			// TODO: POST, GET /:id, PUT, DELETE
		}

		doctors := v1.Group("/doctors")
		doctors.Use(middleware.AuthMiddleware())
		{
			doctors.GET("", func(c *gin.Context) {
				var doctorList []database.Doctor
				database.DB.Find(&doctorList)
				c.JSON(200, doctorList)
			})
		}

		appointments := v1.Group("/appointments")
		appointments.Use(middleware.AuthMiddleware())
		{
			appointments.POST("", middleware.RoleMiddleware("PATIENT", "ADMIN", "STAFF"), appointment.CreateAppointment)
			appointments.PUT("/:id/status", middleware.RoleMiddleware("DOCTOR", "ADMIN", "STAFF"), appointment.UpdateStatus)
		}
	}

	log.Println("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
