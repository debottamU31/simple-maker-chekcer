
# Maker-Checker Message Approval Service

A simplified maker-checker workflow service built with Go, SQLite, and sqlc. The service implements an approval system where "makers" create messages that require "checkers" approval before sending.

## Table of Contents
- [Maker-Checker Message Approval Service](#maker-checker-message-approval-service)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Tech Stack](#tech-stack)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
  - [API Documentation](#api-documentation)
    - [Authentication](#authentication)
    - [Endpoints](#endpoints)
      - [1. Create Message (Maker)](#1-create-message-maker)
      - [2. List Pending Messages (Checker)](#2-list-pending-messages-checker)
      - [3. Approve Message (Checker)](#3-approve-message-checker)
      - [4. Reject Message (Checker)](#4-reject-message-checker)
  - [Enhancements](#enhancements)

## Features
- **Maker-Checker Workflow**
  - Makers create messages (PENDING state)
  - Checkers review and either approve (SENT state) or reject (REJECTED state)
  - Automated message sending upon approval
- **RESTful API** using Chi router
- **Persistent Storage** with SQLite
- **Docker Support** with secure non-root execution
- **Query Management** using sqlc

## Tech Stack
- Go 1.21+
- SQLite
- Chi Router
- sqlc for query generation
- Docker & Docker Compose

## Project Structure
```
.
├── Dockerfile
├── docker-compose.yml
├── schema.sql
├── queries/
│   └── messages.sql
├── sqlc.yaml
├── go.mod
├── go.sum
└── internal/
    ├── db/         # Generated database code
    ├── handler/    # HTTP handlers
    ├── mailer/     # Mailing service
    ├── repository/ # Data access layer
    ├── service/    # Business logic
    └── main.go     # Application entry point
```

## Prerequisites
1. Go 1.21 or higher
2. Docker and Docker Compose
3. sqlc (optional, for query regeneration)
   ```bash
   go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
   ```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Generate database code (optional)**
   ```bash
   sqlc generate
   ```

4. **Initialize database**
   ```bash
   sqlite3 messages.db < schema.sql
   ```

5. **Start the service**
   ```bash
   docker compose up --build
   ```
   The service will be available at `http://localhost:8080`

## API Documentation

### Authentication
All requests require the `X-User-Id` header:
- Maker IDs should start with `maker_`
- Checker IDs should start with `checker_`

### Endpoints

#### 1. Create Message (Maker)
```bash
curl -X POST "http://localhost:8080/messages" \
  -H "Content-Type: application/json" \
  -H "X-User-Id: maker_1" \
  -d '{
    "content": "Hello World",
    "recipient": "john@example.com"
  }'
```

#### 2. List Pending Messages (Checker)
```bash
curl -X GET "http://localhost:8080/messages" \
  -H "X-User-Id: checker_1"
```

#### 3. Approve Message (Checker)
```bash
curl -X POST "http://localhost:8080/messages/{id}/approve" \
  -H "X-User-Id: checker_1"
```

#### 4. Reject Message (Checker)
```bash
curl -X POST "http://localhost:8080/messages/{id}/reject" \
  -H "X-User-Id: checker_1"
```

## Enhancements

- [ ] Implement JWT/OAuth authentication
- [ ] Add support for multiple checker approvals
- [ ] Integrate with external mail service
- [ ] Add configuration management
- [ ] Implement structured logging and metrics
- [ ] Add comprehensive test coverage
