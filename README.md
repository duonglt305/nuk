# NUK

NUK is a Go-based web application framework with built-in authentication, email functionality, and database management.

## Features

- **Authentication System**: User registration, login, profile management, and password recovery
- **Email Service**: SMTP-based email sending with HTML templates
- **Database Support**: MySQL and PostgreSQL support with migrations
- **Dependency Injection**: Using Wire for clean dependency management
- **Configuration Management**: Environment-based configuration using Viper
- **Caching**: Redis integration

## Project Structure

```
├── cmd/                  # Application entry points
│   ├── api/              # API server
│   └── migrate/          # Database migration tool
├── db/                   # Database migrations
│   └── migrations/       # SQL migration files for MySQL and PostgreSQL
├── internal/             # Application core logic
│   ├── auth/             # Authentication module
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
- MySQL or PostgreSQL
- Redis (optional, for caching)

### Setup

1. Clone the repository
2. Initialize the project:

```bash
make init
```

3. Start the development server:

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

## Configuration

NUK uses environment variables for configuration through a `.env` file. An example configuration file (`.env.example`) is provided.

## Email Templates

The application includes HTML email templates with responsive design:

- Base template: `templates/email/base.html`
- Password recovery: `templates/email/passwords/forgot.html`

## License

[Add license information here]