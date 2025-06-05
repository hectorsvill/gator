# [in progress] gator - gator is a RSS feed aggregator
## MVP ( v 0.1.0 )
- [x] Add RSS feeds from across the internet to be collected
- [x] Store the collected posts in a PostgreSQL database
- [x] Follow and unfollow RSS feeds that other users have added
- [ ] View summaries of the aggregated posts in the terminal, with a link to the full post
- [x] integrate a Go application with a PostgreSQL database
- [x] query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go)
- [ ] write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database
## MVP (v 0.2.0)
- [ ] convert to micro service
- [ ] implement gRPC
- [ ] implement a REST API

## How to run
```
# run 
go run . reset
go run . register Jim
go run . addfeed "Hacker News RSS" "https://hnrss.org/newest"
go run . register Andy
go run . addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"
go run . follow "https://hnrss.org/newest"
go run . following
go run . login Dwight
go run . follow "https://hnrss.org/newest"
go run . following
go run . login Jim
go run . following
go run . unfollow "https://hnrss.org/newest"
go run . following
```


## Project setup: [postgresql](https://www.postgresql.org/docs) (mac)
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

## Project setup: [goose](https://github.com/pressly/goose) 
```bash
# install goose
go install github.com/pressly/goose/v3/cmd/goose@latest
#check version
goose -version
```
##### create a users migration 
- create sql file 
```bash  
sql/schema/001_gator.sql
```
- add migragtion code to files
* [001_gator.sql](https://github.com/hectorsvill/gator/blob/main/sql/schema/001_gator.sql)
* [002_feeds.sql](https://github.com/hectorsvill/gator/blob/main/sql/schema/002_feeds.sql)
* [003_feed_follows.sql](https://github.com/hectorsvill/gator/blob/main/sql/schema/003_feed_follows.sql)

### run migrations
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

## Project setup: [sqlc](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

```bash 
# install 
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go get github.com/google/uuid
```
#### touch sqlc.yaml
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

- generate code in internal database
```bash
sqlc generate
```
- import postgresql drivers
```bash
go get github.com/lib/pq
```

