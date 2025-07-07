# FlagChain

A microservice built with **Golang**, **Gin**, and **GORM** to manage feature flags with support for dependencies, audit logging, and automated migrations.

## Features
- **Create** feature flags with optional dependencies.
- **Toggle** flags on/off with validation of dependencies.
- **Cascade disable** dependent flags automatically.
- **Audit log** for all operations (create, toggle, auto-disable) with timestamps, reasons, and actor info. **audit_logs**
- **Cycle detection** to prevent circular dependencies at creation.
- Fully **Dockerized** with `docker-compose`.
- **Automated migrations** using `golang-migrate`.

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://https://github.com/alirezasaharkhiz/FlagChain.git
   cd feature-flags-service
   ```

2. **Environment Variables**
   Copy `.env.example` to `.env` and adjust values:
   ```env
   DB_HOST=db
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=example
   DB_NAME=flagsdb
   SERVER_PORT=:8080
   MIGRATIONS_DIR=./migrations
   ```

3. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```
   - This will start MySQL, run migrations, and launch the API server.

4. **API Endpoints**
   - `POST /flags` – Create a new flag.
   - `GET /flags` – List all flags.
   - `PATCH /flags/:id/toggle` – Toggle a flag on/off.
   - `GET /flags/:id/history` – Retrieve audit log for a flag.

## License
MIT © Alireza Saharkhiz