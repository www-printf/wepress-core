# WePress Core

## Description

Backend service for WePress Smart Printing Project.

## Prerequisites

-   Go (v1.23 or higher)
-   Docker
-   PostgreSQL (via Docker or local installation)
-   Redis (via Docker or local installation)

## Installation

```
git clone https://github.com/www-printf/wepress-core.git
cd wepress-core
```

Create environment file `.env` and configure environment variables

```
cp .env.example .env
```

Setup dependencies

```
make setup
```

## Development

### Run database

Start PostgreSQL database using Docker:

```
make run-db
```

Start Redis using Docker:

```
make run-redis
```

Stop databases:

```
make stop-db
make stop-redis
```

### Run API

```
make run-api
```

Server will listen on `localhost:8080` by default.

Or watch for changes and restart automatically:

```
make watch-api
```

### Database migrations

Generate a new database migration:

```
make migration-create name=<migration_name>
```

Apply migrations:

```
make migration-up
```

Rollback migrations:

```
make migration-down
```

### Protobuf Generation

Generate protocol buffer code:

```
make gen-proto
```

### Testing

Run tests:

```
make test
```

### Documentation

Generate API Documentation:

```
make docs
```
