package appointment

import (
	"clinic-management-api/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAppointmentInput struct {
	DoctorID     uint      `json:"doctor_id" binding:"required"`
	ScheduleDate time.Time `json:"schedule_date" binding:"required"`
	Notes        string    `json:"notes"`
}

type UpdateStatusInput struct {
	Status string `json:"status" binding:"required,oneof=PENDING CONFIRMED IN_PROGRESS COMPLETED CANCELLED"`
}

func CreateAppointment(c *gin.Context) {
	var input CreateAppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	// Find patient by UserID
	var patient database.Patient
	if err := database.DB.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient profile not found"})
		return
	}

	// Check if doctor exists
	var doctor database.Doctor
	if err := database.DB.First(&doctor, input.DoctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	// Business Logic: Check Availability could be added here
	
	appointment := database.Appointment{
		PatientID:    patient.ID,
		DoctorID:     doctor.ID,
		ScheduleDate: input.ScheduleDate,
		Status:       "PENDING",
		Notes:        input.Notes,
	}

	if err := database.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var appointment database.Appointment
	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	appointment.Status = input.Status
	if err := database.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}
