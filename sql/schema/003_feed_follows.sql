-- +goose Up
CREATE TABLE IF NOT EXISTS gator.feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES gator.users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES gator.feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE gator.feed_follows;
