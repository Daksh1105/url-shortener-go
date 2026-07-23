# High-Throughput URL Shortener (Go, PostgreSQL, Redis)

A full-stack, containerized URL shortening microservice built with Go, featuring a Redis cache-aside layer, PostgreSQL persistence, Base62 encoding, and a modern web UI.

## Tech Stack
- **Backend:** Go (Standard Library, Clean Architecture)
- **Database:** PostgreSQL
- **Caching:** Redis
- **Containerization:** Docker & Docker Compose

## Quick Start (Docker Compose)
Clone the repository and spin up the entire stack with a single command:

```bash
docker compose up --build
