# Financial Ledger

An event-driven financial ledger service focused on consistency, idempotency, and distributed architecture.

---

## Overview

This project simulates a real-world financial transaction processing system, ensuring:

- Data consistency with an **immutable ledger pattern** — balance always derived from transactions
- **Idempotency** via idempotency key, preventing duplicate processing
- **Asynchronous and fault-tolerant** processing with RabbitMQ
- Foundation for evolution into a microservices architecture

---

## Tech Stack

- **Go** + **Gin** — HTTP API
- **MongoDB** — primary database
- **RabbitMQ** — messaging with publisher confirms and retry
- **Docker** — containerization

---

## Getting Started

**1. Start MongoDB:**

```bash
docker compose up -d
```

**2. Set environment variables:**

```bash
export MONGO_URI=mongodb://root:rootpassword@localhost:27017/ledger?authSource=admin
export RABBITMQ_URL=amqp://guest:guest@localhost:5672/
export PORT=8080
```

**3. Run the API:**

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`.

---

## Endpoints

Base: `/api/v1`

| Method | Route                            | Description          |
|--------|----------------------------------|----------------------|
| POST   | `/accounts/`                     | Create account       |
| GET    | `/accounts/`                     | List accounts        |
| GET    | `/accounts/:accountId`           | Get account          |
| GET    | `/accounts/:accountId/balance`   | Get balance          |
| GET    | `/accounts/:accountId/statement` | Get statement        |
| POST   | `/transactions/`                 | Create transaction   |
| GET    | `/transactions/`                 | List transactions    |

---

## Roadmap

- [x] Transactions API (credit and debit)
- [x] Idempotency key
- [x] Database persistence
- [x] Event publishing to RabbitMQ
- [ ] Outbox Pattern
- [ ] Publishing worker
- [ ] Event consumer
- [ ] Retry + Dead Letter Queue (DLQ)
- [ ] Split into `ledger-service` and `balance-service`
- [ ] Service communication exclusively via events
- [ ] Structured logging (JSON)
- [ ] Metrics with Prometheus + Grafana
- [ ] Deploy on Kubernetes (EKS)
- [ ] AWS integration (RDS, Redis, SQS)

---

## License

MIT
