# Ledger API

Ledger API is a RESTful service built with Go and MongoDB for managing
financial transactions, accounts, and balances, following clean
architecture principles and running with Docker.

A scalable REST API built with **Go (Golang)** using **MongoDB** as the
database. This project follows clean architecture principles and is
fully containerized with Docker.

------------------------------------------------------------------------

## 🚀 Tech Stack

-   Go (Golang)
-   MongoDB
-   Docker & Docker Compose
-   Clean Architecture
-   RESTful API
-   OpenAPI (Swagger ready)

------------------------------------------------------------------------

## 📁 Project Structure

    ledger-api/
    │
    ├── cmd/
    │   └── api/
    │       └── main.go
    │── docs/
    |
    ├── internal/
    │   ├── database/
    │   ├── handlers/
    │   ├── services/
    │   ├── repositories/
    │   └── models/
    │
    ├── migrations/
    │   └── ...
    │
    ├── Dockerfile
    ├── docker-compose.yml
    └── README.md

------------------------------------------------------------------------

## 🐳 Running with Docker

### 1️⃣ Build and start services

``` bash
docker compose up --build
```

This will start:

-   MongoDB on port `27017`
-   API on port `8080`

------------------------------------------------------------------------

## 🔗 Environment Variables

  Variable    Description
  ----------- -------------------------
  MONGO_URI   Mongo connection string
  PORT        API port (default 8080)

Example:

    MONGO_URI=mongodb://root:rootpassword@mongodb:27017/ledger?authSource=admin

------------------------------------------------------------------------

## 📦 MongoDB Setup

Mongo runs as a service inside Docker.

If using initialization scripts:

    migrations/init.js

These scripts will:

-   Create collections
-   Create indexes
-   Seed initial data

------------------------------------------------------------------------

## 🧠 Migrations Strategy

Since MongoDB is schema-less:

-   Indexes are created on startup
-   Structural changes are handled via Go migrations
-   Migration versions can be stored in a `migrations` collection

Example (Go migration):

``` go
collection.Indexes().CreateOne(ctx, mongo.IndexModel{
    Keys: bson.D{{Key: "email", Value: 1}},
    Options: options.Index().SetUnique(true),
})
```

------------------------------------------------------------------------

## 📖 API Documentation

You can expose OpenAPI/Swagger documentation at:

    /swagger/index.html

(If Swagger is configured in the project)

------------------------------------------------------------------------

## 🛠 Development

### Run locally without Docker

1.  Start MongoDB
2.  Export environment variables
3.  Run:

``` bash
go run cmd/api/main.go
```

------------------------------------------------------------------------

## 📌 Production Considerations

-   Use a managed MongoDB (Atlas / DocumentDB)
-   Enable authentication and TLS
-   Add structured logging
-   Add health checks
-   Add graceful shutdown
-   Add CI/CD pipeline

------------------------------------------------------------------------

## 🧪 Testing

Run:

``` bash
go test ./...
```

------------------------------------------------------------------------

## 📜 License

MIT
