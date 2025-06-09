-- +goose Up
CREATE SCHEMA IF NOT EXISTS gator;

CREATE TABLE gator.users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE gator.users;
DROP SCHEMA gator;

