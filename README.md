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

# postgresql (mac)
#### install 
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

#### create table 
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




