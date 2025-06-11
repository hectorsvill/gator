-- name: CreatePost :one
INSERT INTO gator.posts 
(id, created_at, updated_at, title, description, published_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM gator.posts
WHERE user_id = $1 ORDER BY created_at DESC
LIMIT $2;

-- -- name: DeleteFeeds :exec
-- DELETE FROM gator.feeds;

-- -- name: GetAllFeedNames :many
-- SELECT name FROM gator.feeds;

-- -- name: GetFeedByURL :one
-- SELECT * FROM gator.feeds
-- WHERE url = $1;

-- -- name: MarkFeedFetched :one
-- UPDATE gator.feeds 
-- SET last_fetched_at = NOW(), 
-- updated_at = NOW()
-- WHERE id = $1
-- RETURNING *;

-- -- name: GetNextFeedToFetch :one
-- SELECT * from gator.feeds
-- ORDER BY last_fetched_at ASC NULLS FIRST
-- LIMIT 1;