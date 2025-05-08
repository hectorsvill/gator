-- +goose Up
CREATE SCHEMA gator;

CREATE TABLE gator.users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE gator.feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    url VARCHAR(50) NOT NULL UNIQUE,
    user_id UUID NOT NULL UNIQUE REFERENCES gator.users (id) ON DELETE CASCADE
);

CREATE TABLE gator.feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE gator.users;
DROP TABLE gator.feeds;
DROP TABLE gator.feed_follows;
DROP SCHEMA gator;