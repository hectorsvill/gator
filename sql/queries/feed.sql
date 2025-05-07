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

-- name: GetFeed :one
SELECT * FROM gator.feeds
WHERE url = $1 LIMIT 1;

-- name: DeleteFeeds :exec
DELETE FROM gator.feeds;

-- name: GetAllFeedNames :many
SELECT name FROM gator.feeds;
