# Vacantr

Telegram bot for tracking job vacancies: save, filter, and receive push notifications in real-time.

## Stack

- Go (Golang)
- PostgreSQL
- Redis
- Kafka
- Docker + Docker Compose

## Quick start

```bash
# 1. CLone repository
git clone https://github.com/твой-юзернейм/vacantr.git
cd vacantr

# 2. Copy environment variables
cp .env.example .env

# 3. Build and run all services
docker compose up --build
```

## Project structure

**Follows Clean Architecture principles:**
    
    cmd/
        vacantr/     # Telegram bot entrypoint
        worker/      # Kafka consumer worker
    config/        # Configuration loaders
    internal/      # Business logic
        adapter/     # External interfaces (db, redis, kafka, telegram)
        core/        # Domain models
        usecase/     # Application logic
    migrations/    # SQL schema
    pkg/           # Reusable helpers/utilities

## ENV
    
**TELEGRAM_TOKEN** - Telegram Bot API token
**DB_DSN** - Postgres DSN
**REDIS_ADDR** - Redis connection string
**KAFKA_ADDR** - Kafka broker address

## Bot's command

**/start** - Initialize the bot
**/vacancies** - Show the latest vacancies
**/setfilter** - Set filter keywords (e.g., golang, middle)
**/subscribe** - Enable automatic vacancy notifications
**/unsubscribe** - Stop receiving notifications

## Migrations

Run SQL schema in the Postgres container:

```bash
docker compose exec postgres psql -U tecs -d vacantr -f /migrations/schema.sql
```

## Contacts for help

**Telegram:** @oxtecs