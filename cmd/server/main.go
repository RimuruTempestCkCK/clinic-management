package main

import (
	"clinic-management-api/internal/appointment"
	"clinic-management-api/internal/auth"
	"clinic-management-api/internal/database"
	"clinic-management-api/internal/doctor"
	"clinic-management-api/internal/medical_record"
	"clinic-management-api/internal/medicine"
	"clinic-management-api/internal/middleware"
	"clinic-management-api/internal/patient"
	"clinic-management-api/internal/prescription"
	"log"

	_ "clinic-management-api/docs" // Swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Clinic Management System API
// @version 1.0
// @description REST API for Clinic Management System.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load()
	database.ConnectDB()

	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", auth.Register)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/refresh", auth.Refresh)
			authGroup.POST("/logout", auth.Logout)
		}

		patients := v1.Group("/patients")
		patients.Use(middleware.AuthMiddleware())
		{
			patients.GET("", patient.GetAll)
			patients.POST("", patient.Create)
			patients.GET("/:id", patient.GetByID)
			patients.PUT("/:id", patient.Update)
			patients.DELETE("/:id", patient.Delete)
		}

		doctors := v1.Group("/doctors")
		doctors.Use(middleware.AuthMiddleware())
		{
			doctors.GET("", doctor.GetAll)
			doctors.POST("", doctor.Create)
			doctors.GET("/:id", doctor.GetByID)
			doctors.PUT("/:id", doctor.Update)
			doctors.DELETE("/:id", doctor.Delete)
		}

		appointments := v1.Group("/appointments")
		appointments.Use(middleware.AuthMiddleware())
		{
			appointments.GET("", appointment.GetAll)
			appointments.POST("", appointment.CreateAppointment)
			appointments.GET("/:id", appointment.GetByID)
			appointments.PUT("/:id", appointment.Update)
			appointments.DELETE("/:id", appointment.Delete)
			appointments.PUT("/:id/status", middleware.RoleMiddleware("DOCTOR", "ADMIN", "STAFF"), appointment.UpdateStatus)
		}

		records := v1.Group("/medical-records")
		records.Use(middleware.AuthMiddleware())
		{
			records.GET("", medical_record.GetAll)
			records.POST("", medical_record.Create)
			records.GET("/:id", medical_record.GetByID)
		}

		prescriptions := v1.Group("/prescriptions")
		prescriptions.Use(middleware.AuthMiddleware())
		{
			prescriptions.GET("", prescription.GetAll)
			prescriptions.POST("", prescription.Create)
			prescriptions.GET("/:id", prescription.GetByID)
		}

		medicines := v1.Group("/medicines")
		medicines.Use(middleware.AuthMiddleware())
		{
			medicines.GET("", medicine.GetAll)
			medicines.POST("", medicine.Create)
			medicines.PUT("/:id", medicine.Update)
			medicines.DELETE("/:id", medicine.Delete)
		}
	}

	log.Println("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
