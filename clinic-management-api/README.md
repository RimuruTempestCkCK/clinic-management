# 🏥 Clinic Management System API

This is the backend API for a complete Clinic Management System. Built using Go, Gin, GORM, and PostgreSQL (Supabase).

## 🚀 Tech Stack

- **Go (Golang)**
- **Gin Framework** - Fast HTTP web framework
- **GORM** - Developer-friendly ORM library
- **PostgreSQL** - Hosted on Supabase (using Connection Pooling)
- **JWT** - Secure stateless authentication
- **Swagger / Swaggo** - Auto-generated interactive API Documentation
- **Bcrypt** - Password hashing

## 📂 Project Structure

```text
clinic-management-api/
├── cmd/
│   └── server/
│       └── main.go       # Application Entrypoint
├── docs/                 # Swagger Documentation Files
├── internal/
│   ├── appointment/      # Business logic & Handlers for Appointments
│   ├── auth/             # Authentication & JWT generation
│   ├── config/           # Configuration loaders
│   ├── database/         # Database connection & Models
│   ├── doctor/           # Handlers for Doctors
│   ├── medical_record/   # Handlers for Medical Records
│   ├── medicine/         # Handlers for Medicines
│   ├── middleware/       # JWT & Role-Based Access Control (RBAC) middlewares
│   ├── patient/          # Handlers for Patients
│   ├── prescription/     # Handlers for Prescriptions
│   └── user/             # Handlers for Users
├── .env.example          # Environment variables template
├── seed_with_schema.sql  # Database Schema & Seed Data (Dummy records)
```

## 👥 Roles & Access Control

The API implements strict RBAC (Role-Based Access Control):
- **ADMIN**: Full access to all resources.
- **DOCTOR**: Access to patients, appointments, medical records, and prescriptions.
- **STAFF**: Can manage appointments and patients.
- **PATIENT**: Can view their own records and book appointments.

## 🛠️ Setup & Installation

### 1. Database Setup (Supabase)
Run the provided SQL script to create the schema and insert dummy data:
1. Open your Supabase Dashboard -> **SQL Editor**.
2. Copy the contents of `seed_with_schema.sql`.
3. Run the script.

### 2. Environment Variables
Rename `.env.example` to `.env` and fill in your connection details:
```env
DATABASE_URL="postgres://postgres.[YOUR_PROJECT_REF]:[PASSWORD]@aws-0-[REGION].pooler.supabase.com:5432/postgres"
JWT_SECRET="your_secret_key"
PORT=8080
```

### 3. Run the API Server
```bash
go run cmd/server/main.go
```

## 📖 API Documentation (Swagger UI)

Once the server is running, you can access the interactive API documentation and test endpoints directly from your browser:
👉 **http://localhost:8080/swagger/index.html**

### How to Authenticate in Swagger
1. Call `POST /api/v1/auth/login` with email `admin@clinic.com` and password `password123`.
2. Copy the token from the response.
3. Scroll to the top of Swagger, click the **Authorize** button.
4. Paste the token exactly as it is (no need to type "Bearer", the system handles it) and click Authorize.

## 🔗 Endpoints Overview

| Feature | Endpoints |
| --- | --- |
| **Auth** | `POST /auth/register`, `POST /auth/login` |
| **Patients** | `GET /patients`, `POST /patients`, `PUT /patients/:id`, `DELETE /patients/:id` |
| **Doctors** | `GET /doctors`, `POST /doctors`, `PUT /doctors/:id`, `DELETE /doctors/:id` |
| **Appointments**| `GET /appointments`, `POST /appointments`, `PUT /appointments/:id/status` |
| **Medical Rec.**| `GET /medical-records`, `POST /medical-records`, `GET /medical-records/:id` |
| **Medicines** | `GET /medicines`, `POST /medicines`, `PUT /medicines/:id`, `DELETE /medicines/:id` |
| **Prescriptions**|`GET /prescriptions`, `POST /prescriptions`, `GET /prescriptions/:id` |
