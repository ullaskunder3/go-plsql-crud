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

### ✅ **3. Implemented CRUD Operations**
- Implemented Create, Read, Update, and Delete operations using GORM.
- Tested the operations with both the database and Go models.

### ✅ **4. Built API Endpoints**
- Set up Gin for RESTful API routing.
- Implemented routes for creating, fetching, updating, and deleting users.
- Tested the API using cURL and Postman.

---

## 📝 Next Steps
- Add authentication middleware (JWT or basic authentication).
- Introduce input validation in service layer.
- Refactor for better separation of concerns (e.g., services for business logic).
- Containerize the app using Docker.

---

## ⚙️ Tech Stack
- **Go** (Golang)
- **PostgreSQL**
- **GORM** ORM
- **Gin** (Web Framework for API)
- **cURL** (API Testing)
- **Docker** (Planned for later)

---

## 📜 How to Run

1. **Clone the Repository**  
```bash
git clone https://github.com/YOUR_USERNAME/go-postgres-crud.git
cd go-postgres-crud
```

2. **Set Up Environment Variables**  
Create a `.env` file based on the example template:
```bash
cp .env.example .env
```
Then, edit the `.env` file with your PostgreSQL credentials.

3. **Run the Application**  
To start the Go application and the Gin server:
```bash
go run cmd/main.go
```
The server will start running on `http://localhost:8080`.

---

## 💡 Notes
- Ensure **PostgreSQL** is running before starting the application.
- Use `psql` to manually inspect the database if needed.
- The app is currently configured to run with no authentication. Plan to add JWT authentication for production use.

---

## 🌍 Testing the API (cURL examples)

### **Create a User**
```bash
curl -X POST "http://localhost:8080/users" \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice", "email": "alice@example.com"}'
```
Response:
```json
{
    "ID": 1,
    "CreatedAt": "2025-04-01T22:12:30.303790226+05:30",
    "UpdatedAt": "2025-04-01T22:12:30.303790226+05:30",
    "DeletedAt": null,
    "Name": "Alice",
    "Email": "alice@example.com"
}
```

### **Get All Users**
```bash
curl -X GET "http://localhost:8080/users"
```
Response:
```json
[
    {
        "ID": 1,
        "Name": "Alice",
        "Email": "alice@example.com"
    }
]
```

### **Get a User by ID**
```bash
curl -X GET "http://localhost:8080/users/1"
```
Response:
```json
{
    "ID": 1,
    "Name": "Alice",
    "Email": "alice@example.com"
}
```

### **Update a User**
```bash
curl -X PUT "http://localhost:8080/users/1" \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Updated", "email": "alice_updated@example.com"}'
```
Response:
```json
{
    "ID": 1,
    "Name": "Alice Updated",
    "Email": "alice_updated@example.com"
}
```

### **Delete a User**
```bash
curl -X DELETE "http://localhost:8080/users/1"
```
Response:
```json
{
    "message": "User deleted successfully"
}
```

---

Stay tuned for more updates! 🚀