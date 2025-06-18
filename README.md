# NUK

NUK is a Go-based web application framework with built-in authentication, email functionality, and database management. It follows clean architecture principles with dependency injection.

## Features

- **Authentication System**: User registration, login, profile management, and password recovery
- **Email Service**: SMTP-based email sending with HTML templates
- **Database Support**: MySQL and PostgreSQL support with migrations
- **Dependency Injection**: Using Wire for clean dependency management
- **Configuration Management**: Environment-based configuration using Viper
- **Caching**: Redis integration
- **Hot Reload**: Development with Air for auto-reloading
- **Docker Support**: Complete Docker Compose setup for development

## Project Structure

```
├── cmd/                  # Application entry points
│   ├── api/              # API server
│   └── migrate/          # Database migration tool
├── data/                 # Data directories for Docker volumes
│   ├── postgres/         # PostgreSQL data
│   └── redis/            # Redis data
├── db/                   # Database migrations
│   └── migrations/       # SQL migration files
│       ├── mysql/        # MySQL-specific migrations
│       └── pgsql/        # PostgreSQL-specific migrations
├── internal/             # Application core logic
│   ├── auth/             # Authentication module with clean architecture
│   │   ├── application/  # Application services and DTOs
│   │   ├── domain/       # Business domain entities and repositories
│   │   ├── infrastructure/ # Infrastructure implementations
│   │   └── presentation/ # HTTP handlers and middleware
│   └── ...
├── pkg/                  # Reusable packages
│   ├── cache/            # Caching utilities
│   ├── db/               # Database utilities
│   ├── email/            # Email services
│   ├── http/             # HTTP utilities
│   ├── jwt/              # JWT utilities
│   └── utils/            # General utilities
└── templates/            # HTML templates
    └── email/            # Email templates
```

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose (recommended for development)
- MySQL or PostgreSQL
- Redis (for caching)
- Mailpit (for email testing)

### Setup (Local Development)

#### Option 1: Using Docker (Recommended)

1. Clone the repository
2. Initialize the project:

```bash
make init
```

3. Start the development environment with Docker Compose:

```bash
docker-compose up -d
```

This will start:
- NUK API service with hot reloading
- PostgreSQL database
- Redis cache
- Mailpit for email testing (accessible at http://localhost:8025)

#### Option 2: Manual Setup

1. Clone the repository
2. Initialize the project:

```bash
make init
```

3. Start the development server with hot reloading:

```bash
make dev
```

### Database Migrations

Create a new migration:

```bash
make migration name=your_migration_name
```

Run migrations:

```bash
make migrate
```

### Configuration

NUK uses environment variables for configuration through a `.env` file. An example configuration file (`.env.example`) is provided with the following settings:

```bash
PORT=8080                        # Application port
JWT_SECRET=secret                # Secret for JWT signing
JWT_ACCESS_TOKEN_LIFETIME=3600   # Access token lifetime in seconds
JWT_REFRESH_TOKEN_LIFETIME=86400 # Refresh token lifetime in seconds
SF_WORKER=1                      # Snowflake ID worker identifier
DB_DRIVER=pgsql                  # Database driver (pgsql or mysql)
DB_URL=                          # Database connection URL
CACHE_DRIVER=redis               # Cache driver
CACHE_URL=                       # Cache connection URL
MAIL_HOST=                       # SMTP host
MAIL_PORT=                       # SMTP port
MAIL_USER=                       # SMTP username
MAIL_PASSWORD=                   # SMTP password
MAIL_FROM_ADDRESS=               # Sender email address
MAIL_FROM_NAME=                  # Sender name
```

## Email Templates

The application includes responsive HTML email templates:

- Base template: `templates/email/base.html`
- Partials: `templates/email/partials/`
- Password recovery: `templates/email/passwords/forgot.html`

## Development Tools

NUK uses several development tools to streamline the development process:

- **Air**: For hot reloading during development
- **Wire**: For dependency injection code generation
- **Golang Migrate**: For database migration management

## License

MIT License

Copyright (c) 2025 duonglt.net

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.