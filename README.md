# Go-Gator

Go-Gator is a command-line RSS aggregator built with Go and PostgreSQL.

It allows users to subscribe to RSS feeds, follow feeds created by other users, continuously aggregate new content, and browse posts directly from the terminal.

## Features

- User registration and login
- RSS feed management
- Follow and unfollow feeds
- Background feed aggregation
- Automatic post collection and storage
- Browse posts from followed feeds
- PostgreSQL-backed persistence
- Database migrations with Goose
- Type-safe SQL queries with SQLC

## Requirements

Before installing Go-Gator, ensure you have:

- Go 1.24 or newer
- PostgreSQL

Verify your installation:

```bash
go version
psql --version
```

## Installation

Install the latest version directly from GitHub:

```bash
go install github.com/shubh1855/go-go-gator@latest
```

Ensure your Go binary directory is included in your PATH.

## Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE go-gator;
```

Run the database migrations:

```bash
goose postgres "<connection_string>" up
```

Example:

```bash
goose postgres "postgres://postgres:password@localhost:5432/go-gator" up
```

## Configuration

Create a configuration file at:

```text
~/.gatorconfig.json
```

Example:

```json
{
  "db_url": "postgres://postgres:password@localhost:5432/go-gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace the connection string with your PostgreSQL credentials.

## Usage

### Register a User

```bash
go-gator register alice
```

### Login

```bash
go-gator login alice
```

### List Users

```bash
go-gator users
```

### Add a Feed

```bash
go-gator addfeed "Hacker News" https://news.ycombinator.com/rss
```

### View All Feeds

```bash
go-gator feeds
```

### Follow a Feed

```bash
go-gator follow https://news.ycombinator.com/rss
```

### View Followed Feeds

```bash
go-gator following
```

### Unfollow a Feed

```bash
go-gator unfollow https://news.ycombinator.com/rss
```

### Start the Aggregator

Fetch new posts from subscribed feeds at a fixed interval:

```bash
go-gator agg 30s
```

Examples:

```bash
go-gator agg 10s
go-gator agg 1m
go-gator agg 5m
```

### Browse Posts

Show the latest posts from feeds you follow:

```bash
go-gator browse
```

Specify a custom limit:

```bash
go-gator browse 10
```

## Architecture

### Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Migrations:** Goose
- **Query Generation:** SQLC
- **Feed Format:** RSS/XML

### Components

- **CLI Layer** – command handling and user interaction
- **Database Layer** – SQLC-generated query code
- **RSS Layer** – feed fetching and parsing
- **Aggregator** – background feed collection service

## Development

Generate SQLC code:

```bash
sqlc generate
```

Run migrations:

```bash
goose postgres "<connection_string>" up
```

Rollback the latest migration:

```bash
goose postgres "<connection_string>" down
```

Run the application locally:

```bash
go run .
```

## License

MIT License
