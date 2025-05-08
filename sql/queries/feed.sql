-- name: CreateFeed :one
INSERT INTO gator.feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedNameUrlUser :many
SELECT name, url, user_id FROM gator.feeds;

-- name: DeleteFeeds :exec
DELETE FROM gator.feeds;

-- name: GetAllFeedNames :many
SELECT name FROM gator.feeds;

-- name: GetFeedByURL :one
SELECT * FROM gator.feeds
WHERE url = $1;

