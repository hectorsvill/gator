-- +goose Up
CREATE TABLE IF NOT EXISTS gator.posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR NOT NULL,
    published_at VARCHAR NOT NULL,
    feed_id VARCHAR NOT NULL,
    user_id UUID NOT NULL REFERENCES gator.users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE gator.posts;
