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
   git clone https://github.com/alirezasaharkhiz/FlagChain.git
   cd FlagChain
   ```

2. **Environment Variables**
   Copy `.env.example` to `.env` and adjust values:
   ```env
   DB_HOST=db
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=
   DB_NAME=flag_chain_db
   SERVER_PORT=:8080
   MIGRATIONS_DIR=./migrations
   ```

3. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```
   - This will start MySQL, run migrations, and launch the API server.

4. **API Endpoints**
   - `GET /api/ping` – Ping endpoint, returns `"pong"`.
   - `POST /api/flags` – Create a new flag (with or without dependencies).
   - `GET /api/flags` – List all flags and dependencies.
   - `PUT /api/flags/:id/toggle` – Toggle a flag on/off (If a dependency flag is disabled, any flags depending on it will also be automatically disabled).
   - `GET /api/flags/:id/history` – Retrieve audit log for a flag.
   - `POST /api/flags/:id/dependencies` – Add a dependency to a flag, with checking circular dependency.

## License
MIT © Alireza Saharkhiz
