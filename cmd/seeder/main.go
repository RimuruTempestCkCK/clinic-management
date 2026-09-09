package main

import (
	"clinic-management-api/internal/database"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load env
	_ = godotenv.Load()

	// Connect DB and migrate
	database.ConnectDB()

	fmt.Println("🚀 Starting Database Seeder...")

	gofakeit.Seed(0) // Random seed

	// 1. Create Default Admin
	hashedAdmin, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	adminUser := database.User{
		Email:    "admin@clinic.com",
		Password: string(hashedAdmin),
		Role:     "ADMIN",
	}
	database.DB.FirstOrCreate(&adminUser, database.User{Email: "admin@clinic.com"})

	// 2. Create Specializations
	specs := []string{"Cardiology", "Neurology", "Pediatrics", "Dermatology", "General Practice"}
	var dbSpecs []database.Specialization
	for _, s := range specs {
		spec := database.Specialization{Name: s}
		database.DB.FirstOrCreate(&spec, database.Specialization{Name: s})
		dbSpecs = append(dbSpecs, spec)
	}

	// 3. Create Doctors (20 Doctors)
	fmt.Println("👨‍⚕️ Seeding 20 Doctors...")
	var doctors []database.Doctor
	for i := 0; i < 20; i++ {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("doctor123"), bcrypt.DefaultCost)
		user := database.User{
			Email:    gofakeit.Email(),
			Password: string(hashed),
			Role:     "DOCTOR",
		}
		database.DB.Create(&user)

		doc := database.Doctor{
			UserID:           user.ID,
			Name:             "Dr. " + gofakeit.Name(),
			SpecializationID: dbSpecs[gofakeit.Number(0, len(dbSpecs)-1)].ID,
			Phone:            gofakeit.Phone(),
			Bio:              gofakeit.Sentence(10),
		}
		database.DB.Create(&doc)
		doctors = append(doctors, doc)
	}

	// 4. Create Patients (100 Patients)
	fmt.Println("🤒 Seeding 100 Patients...")
	var patients []database.Patient
	for i := 0; i < 100; i++ {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("patient123"), bcrypt.DefaultCost)
		user := database.User{
			Email:    gofakeit.Email(),
			Password: string(hashed),
			Role:     "PATIENT",
		}
		database.DB.Create(&user)

		gender := "Male"
		if gofakeit.Bool() {
			gender = "Female"
		}

		patient := database.Patient{
			UserID:  user.ID,
			Name:    gofakeit.Name(),
			DOB:     gofakeit.Date(),
			Gender:  gender,
			Phone:   gofakeit.Phone(),
			Address: gofakeit.Address().Address,
		}
		database.DB.Create(&patient)
		patients = append(patients, patient)
	}

	// 5. Create Medicines (50 Medicines)
	fmt.Println("💊 Seeding 50 Medicines...")
	var medicines []database.Medicine
	for i := 0; i < 50; i++ {
		med := database.Medicine{
			Name:        gofakeit.Word() + " " + gofakeit.Word(),
			Description: gofakeit.Sentence(5),
			Price:       gofakeit.Price(10, 500),
			Stock:       gofakeit.Number(10, 1000),
		}
		database.DB.Create(&med)
		medicines = append(medicines, med)
	}

	// 6. Create Appointments, Medical Records, and Prescriptions (500 total)
	fmt.Println("📅 Seeding 500 Appointments & Records...")
	statuses := []string{"PENDING", "CONFIRMED", "IN_PROGRESS", "COMPLETED", "CANCELLED"}
	
	for i := 0; i < 500; i++ {
		patient := patients[gofakeit.Number(0, len(patients)-1)]
		doctor := doctors[gofakeit.Number(0, len(doctors)-1)]
		status := statuses[gofakeit.Number(0, len(statuses)-1)]

		appt := database.Appointment{
			PatientID:    patient.ID,
			DoctorID:     doctor.ID,
			ScheduleDate: gofakeit.FutureDate(),
			Status:       status,
			Notes:        gofakeit.Sentence(5),
		}
		database.DB.Create(&appt)

		// If completed, add Medical Record
		if status == "COMPLETED" {
			record := database.MedicalRecord{
				AppointmentID: appt.ID,
				PatientID:     patient.ID,
				DoctorID:      doctor.ID,
				Diagnosis:     gofakeit.Sentence(3),
				Treatment:     gofakeit.Sentence(5),
				Notes:         gofakeit.Sentence(4),
			}
			database.DB.Create(&record)

			// 50% chance to have a prescription
			if gofakeit.Bool() {
				prescription := database.Prescription{
					MedicalRecordID: record.ID,
					Notes:           "Take after meals",
				}
				database.DB.Create(&prescription)

				// Add 1-3 random medicines
				numItems := gofakeit.Number(1, 3)
				for j := 0; j < numItems; j++ {
					med := medicines[gofakeit.Number(0, len(medicines)-1)]
					item := database.PrescriptionItem{
						PrescriptionID: prescription.ID,
						MedicineID:     med.ID,
						Quantity:       gofakeit.Number(1, 20),
						Dosage:         "2x a day",
					}
					database.DB.Create(&item)
				}
			}
		}
	}

	fmt.Println("✅ Database Seeding Completed Successfully!")
}
