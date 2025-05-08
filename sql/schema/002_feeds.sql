-- +goose Up
CREATE TABLE gator.feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    url VARCHAR(50) NOT NULL UNIQUE,
    user_id UUID NOT NULL UNIQUE REFERENCES gator.users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE gator.feeds;
