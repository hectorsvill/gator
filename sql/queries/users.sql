-- name: CreateUser :one
INSERT INTO gator.users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM gator.users
WHERE name = $1 LIMIT 1;

-- name: DeleteUsers :exec
DELETE FROM gator.users;

-- name: GetAllUsers :many
SELECT * FROM gator.users;

