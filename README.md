# Orders API

Orders API is a simple service built using Go (without any framework) to manage orders. This API supports CRUD operations (Create, Read, Update, Delete) and can be deployed using Docker.

## 🚀 Features
- **CRUD Order**: Create, read, update, and delete orders.
- **Dockerized**: Can be run inside a Docker container.

---
## 📌 Project Structure
```
/orders-api
├── handlers/order_handler.go    # API order handler
├── models/order.go              # Order data model
├── routes/routes.go             # API routing
├── main.go                      # Application entry point
├── go.mod                       # Go dependencies
├── Dockerfile                   # Docker setup
└── README.md                    # Project documentation
```

---
## 🛠️ How to Run

### **1️⃣ Run Locally**
Make sure Go is installed, then run:
```sh
go run main.go
```
The API will be available at `http://localhost:8080`

### **2️⃣ Run with Docker**
```sh
docker build -t orders-api .
docker run -d -p 8080:8080 orders-api
```
Check the API at `http://localhost:8080`

---
## 📦 API Endpoints
| Method | Endpoint      | Description        |
|--------|-------------|----------------|
| GET    | `/orders`   | Retrieve all orders |
| GET    | `/orders/{id}` | Retrieve order by ID |
| POST   | `/orders`   | Create a new order |
| PUT    | `/orders/{id}` | Update an order |
| DELETE | `/orders/{id}` | Delete an order |

---
## 📜 License
This project is licensed under the MIT License. Feel free to use and modify it as needed!

