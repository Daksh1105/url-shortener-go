#  High-Throughput URL Shortener

A high-performance, containerized full-stack URL shortening microservice built with **Go**, **PostgreSQL**, and **Redis**. Built following **Clean Architecture** principles, featuring a **Cache-Aside** design pattern for fast reads and **Base62 encoding** for short path generation.

---

##  Tech Stack

* **Backend:** Go (Standard Library, Clean Architecture)
* **Database:** PostgreSQL (Persistent Storage)
* **Caching:** Redis (Cache-Aside Layer)
* **Containerization:** Docker & Docker Compose (Multi-stage build)
* **Frontend:** HTML5, CSS3, Vanilla JavaScript

---

##  Key Architectural Features

* **Clean Architecture:** Strict decoupling between HTTP handlers, business logic, and database layer (`cmd/`, `internal/`, `pkg/`).
* **Cache-Aside Pattern:** High-frequency redirects hit Redis memory first for sub-millisecond retrieval. DB queries on cache misses backfill Redis automatically.
* **Base62 Encoding:** Encodes auto-incrementing integer IDs into compact, URL-safe strings (e.g., `http://localhost:8080/aX9z`).
* **Multi-Stage Docker Container:** Minimal Alpine runtime container for small image size and rapid deployments.

---

##  Quick Start

### Option 1: Docker Compose (Recommended)

Run the entire application stack (Go Server + Postgres + Redis) with a single command:

```bash
# 1. Clone repository
git clone [https://github.com/Dakshg1105/url-shortener-go.git](https://github.com/Dakshg1105/url-shortener-go.git)
cd url-shortener-go

# 2. Launch container stack
docker compose up --build

##  Access the Application

Open the web UI in your browser:

```text
http://localhost:8080
```

---

##  Option 2: Pull Pre-built Image from Docker Hub

```bash
docker pull docker.io/dakshg1105/url-shortener-go:v1.0
```

---

##  API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST** | `/api/shorten` | Accepts `{ "url": "https://example.com" }` and returns a shortened URL |
| **GET** | `/{shortCode}` | Redirects (`302 Found`) to the original URL |
| **GET** | `/` | Serves the web UI |

---

##  Repository Structure

```text
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/             # Private application logic (Handlers, Services, Storage)
├── web/                  # Frontend UI files (HTML, CSS, JS)
├── Dockerfile            # Multi-stage container definition
├── docker-compose.yml    # Service orchestration (Go, Postgres, Redis)
├── go.mod                # Module definitions
└── README.md
```

---

##  Push to GitHub

Create or open `README.md` in your project folder, paste the contents into it, and run:

```bash
git add README.md
git commit -m "docs: add comprehensive README.md"
git push
```
