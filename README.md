# 🏥 Clinic Management System

A full-stack Clinic Management System monorepo, featuring a highly-performant **Golang** backend API and a beautiful, modern **SvelteKit** web frontend utilizing the Adminator template.

This system is designed to handle daily clinic operations including user management, doctor scheduling, patient tracking, appointments, medical records, and prescriptions.

---

## 📂 Project Structure

This repository is structured as a monorepo containing two main projects:

- `clinic-management-api/` - The Backend REST API (Golang, Gin, GORM, PostgreSQL)
- `clinic-management-web/` - The Frontend Web Application (SvelteKit 5, Adminator UI, Tailwind)

---

## ⚙️ Backend (API)

The backend is a robust RESTful API built with Go, offering strict Role-Based Access Control (RBAC) and comprehensive Swagger documentation.

### Tech Stack
- **Go (Golang)**
- **Gin Framework** - Fast HTTP web framework
- **GORM** - Developer-friendly ORM library
- **PostgreSQL / Supabase** - Relational database
- **JWT & Bcrypt** - Secure stateless authentication
- **Swagger / Swaggo** - Auto-generated API documentation

### Key Features
- **RBAC (Role-Based Access Control):** Supports `ADMIN`, `DOCTOR`, `STAFF`, and `PATIENT` roles.
- **Appointments Management:** Track status (Pending -> Confirmed -> In Progress -> Completed -> Cancelled).
- **CORS Enabled:** Fully integrated with the frontend out-of-the-box.
- **Database Seeding:** Pre-configured SQL scripts to seed the database with dummy data.

### How to Run the API
1. Navigate to the API directory:
   ```bash
   cd clinic-management-api
   ```
2. Rename `.env.example` to `.env` and fill in your connection details (e.g., Supabase Postgres URL).
3. Run the database seed script `seed_with_schema.sql` in your database.
4. Run the server:
   ```bash
   go run cmd/server/main.go
   ```
5. View API Documentation via Swagger: `http://localhost:8080/swagger/index.html`

---

## 🖥️ Frontend (Web)

The frontend is a fast, reactive web application built with the latest Svelte 5 (Runes Mode), leveraging the elegant Adminator dashboard template.

### Tech Stack
- **Svelte 5 / SvelteKit** - Next-generation reactive web framework
- **Adminator Template** - Professional dashboard UI shell
- **Tailwind CSS** - Utility-first CSS framework
- **Vite** - Lightning-fast frontend tooling

### Key Features
- **Beautiful Auth Flow:** Clinic-themed custom login page.
- **Dynamic Dashboard:** Real-time metrics, recent appointments, and doctor rosters fetched directly from the backend.
- **Responsive Layout:** Reusable Svelte components (`Sidebar`, `Header`, `Footer`) perfectly integrated with the template's layout system.

### How to Run the Web App
1. Navigate to the Web directory:
   ```bash
   cd clinic-management-web
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ```
4. Access the application in your browser: `http://localhost:5173`
   *(Login with `admin@clinic.com` / `password123` based on the database seeder).*

---

## 🚀 Quick Start (Running Both)

To run the complete system locally, you'll need two terminal windows:

**Terminal 1 (Backend):**
```bash
cd clinic-management-api
go run cmd/server/main.go
```

**Terminal 2 (Frontend):**
```bash
cd clinic-management-web
npm run dev
```

Visit `http://localhost:5173` to interact with the full system!
