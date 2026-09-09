package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);not null"` // ADMIN, DOCTOR, STAFF, PATIENT
}

type Patient struct {
	gorm.Model
	UserID  uint
	User    User
	Name    string
	DOB     time.Time
	Gender  string
	Phone   string
	Address string
}

type Specialization struct {
	gorm.Model
	Name string
}

type Doctor struct {
	gorm.Model
	UserID           uint
	User             User
	Name             string
	SpecializationID uint
	Specialization   Specialization
	Phone            string
	Bio              string
}

type Appointment struct {
	gorm.Model
	PatientID    uint
	Patient      Patient
	DoctorID     uint
	Doctor       Doctor
	ScheduleDate time.Time
	Status       string `gorm:"type:varchar(20);default:'PENDING'"` // PENDING, CONFIRMED, IN_PROGRESS, COMPLETED, CANCELLED
	Notes        string
}

type MedicalRecord struct {
	gorm.Model
	AppointmentID uint
	Appointment   Appointment
	PatientID     uint
	Patient       Patient
	DoctorID      uint
	Doctor        Doctor
	Diagnosis     string
	Treatment     string
	Notes         string
}

type Medicine struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
}

type Prescription struct {
	gorm.Model
	MedicalRecordID uint
	MedicalRecord   MedicalRecord
	Notes           string
}

type PrescriptionItem struct {
	gorm.Model
	PrescriptionID uint
	Prescription   Prescription
	MedicineID     uint
	Medicine       Medicine
	Quantity       int
	Dosage         string
}
