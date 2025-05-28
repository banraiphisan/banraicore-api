# Banraicore API

Golang-based RESTful API service for Banrai Phisan, built using Clean Architecture.

## 🔧 Tech Stack

- **Language:** Go (Golang)
- **Architecture:** Clean Architecture (Domain → Usecase → Infrastructure)
- **Database:** (Not explicitly found; assumed PostgreSQL or compatible)
- **API Docs:** OpenAPI 3 (see `docs/openapi.yaml`)
- **Deployment:** Docker, GitHub Actions
- **Cache:** In-Memory / Redis

## 📁 Project Structure

```plaintext
cmd/              → Application entry point (main.go)
config/           → Application config (env, port, etc.)
internal/
  ├─ domain/      → Data models and business entities
  ├─ usecase/     → Feature-based logic (auth, user, room, etc.)
  └─ app/         → App initialization and bootstrapping
pkg/              → Shared utilities (e.g., caching)
docs/             → OpenAPI docs and specs
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Docker (optional)

### Build and Run

```bash
# Build with make
make build

# Run the service
make run
```

### Run via Docker

```bash
docker build -t banraicore-api .
docker run -p 8080:8080 banraicore-api
```

## 📚 API Documentation

- Visit `http://localhost:8080/swagger` (or refer to `docs/openapi.yaml`)

## 📂 Features

- ✅ Auth (JWT-based)
- ✅ User Management
- ✅ Room and Reservation System
- ✅ Short URL service
- ✅ Role-based permission system

## 🧪 Testing

```bash
make test
```

## 🛠 Deployment

- GitHub Actions configured in `.github/workflows/deploy.yml`

---

## 📝 License

This project is licensed under the MIT License.
