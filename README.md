## Features
- Pagination
- Tagging
- Adding comments
- RSS feed similuation
- Full text search ability

## Requirements
- Install [Golang](https://go.dev/doc/install)
- Install [Postgres](https://www.postgresql.org/docs/current/install-binaries.html)

## Env Variables
- `BLOG_HOST`: Application host.
- `BLOG_PORT`: Application port
- `BLOG_USER`: Postrgres user
- `BLOG_DBNAME`: Postgres database name
- `BLOG_DBPASSWORD`: Postgres password

## To Run
- Clone repo and cd into repo.
- Set environment variables.
- Seed database: `cat setup.sql | sudo -u postgres psql`
- Run: `go run .`