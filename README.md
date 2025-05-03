# [in progress] gator - gator is a RSS feed aggregator
## MVP
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post
## 
- integrate a Go application with a PostgreSQL database
- query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go)
- write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

# [postgresql](https://www.postgresql.org/docs) (mac)
- install 
```bash
# install
brew install postgresql@15

# start service
brew services start postgresql@15

# check verison
psql --version

# enter psql shell
psql postgres
```

- create table 
```bash
# enter psql shell
psql postgres
# create a db
postgres=# CREATE DATABASE gator;
# enter db
postgres=# \c gator
# check db version 
gator=# SELECT version();
```

#### goose 
```bash
# install goose
go install github.com/pressly/goose/v3/cmd/goose@latest
#check version
goose -version
```
#### create a users migration 
- create sql file 
```bash  
sql/schema/001_users.sql
```
- create uuid extention for 3p lib uuid-ossp
```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
- add migragtion code to file
```sql
-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;
```
- connection string
```bash
psql "postgres://<user name>:@localhost:5432/gator"
```
- run migrations
```
goose postgres <connection_string> up
goose postgres <connection_string> down
```
- view 
```bash
psql gator
\dt
```

- update config with
```json
{"db_url":"postgres://<user_name>:@localhost:5432/gator?sslmode=disable"}
```

# [sqlc](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

```bash 
# install 
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go get github.com/google/uuid
```
# touch sqlc.yaml
```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```
- add sql query to sql/queries
```sql
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;
```
- generate code in internal database
```bash
sqlc generate
```
- import postgresql drivers
```bash
go get github.com/lib/pq
```

