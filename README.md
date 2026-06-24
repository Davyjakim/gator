# Gator

Gator is a command-line RSS aggregator built using Go, PostgreSQL, Goose, and sqlc. It allows you to register accounts, follow your favorite RSS feeds, and browse posts directly from your terminal.

## Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Database Migrations:** Goose
- **Type-Safe SQL Generation:** sqlc

---

## Prerequisites & Setup

Before running Gator, ensure you have Go and PostgreSQL installed on your machine.

### 1. Database Configuration
Create a PostgreSQL database and run your database migrations using Goose:
```bash
go install [github.com/pressly/goose/v3/cmd/goose@latest](https://github.com/pressly/goose/v3/cmd/goose@latest)
goose -dir sql/schema postgres "postgres://username:password@localhost:5432/gator_db" up