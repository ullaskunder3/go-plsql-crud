# Go + PostgreSQL CRUD

## 🚀 Project Overview
This project is a simple CRUD (Create, Read, Update, Delete) application built using **Go** with **PostgreSQL** as the database. We are using **GORM** as the ORM (Object Relational Mapper) to handle database interactions.

## 📌 Journey So Far
### ✅ **1. Setting Up PostgreSQL**
- Installed and configured PostgreSQL on Fedora.
- Initialized the database and set up authentication.
- Resolved authentication errors by modifying `pg_hba.conf`.

### ✅ **2. Connecting Go with PostgreSQL using GORM**
- Created a `.env` file to store database credentials securely.
- Configured a Go application to connect with PostgreSQL.
- Successfully established a connection using GORM.

## 📝 Next Steps
- Define models and migrations.
- Implement CRUD operations.
- Build API endpoints using Go's `net/http` package.
- Add authentication and validation.

## ⚙️ Tech Stack
- **Go** (Golang)
- **PostgreSQL**
- **GORM** ORM
- **Docker** (Planned for later)

## 📜 How to Run
```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/go-postgres-crud.git
cd go-postgres-crud

# Set up environment variables
cp .env.example .env

# Run the application
go run main.go
```

## 💡 Notes
- Ensure PostgreSQL is running before starting the application.
- Use `psql` to manually inspect the database if needed.

---
Stay tuned for more updates! 🚀