-- +goose Up
CREATE TABLE IF NOT EXISTS gator.feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL REFERENCES gator.users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE gator.feeds;




