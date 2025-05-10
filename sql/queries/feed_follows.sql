-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO gator.feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT
    inserted_feed_follow.*,
    gator.feeds.name AS feed_name,
    gator.users.name AS user_name
FROM inserted_feed_follow
INNER JOIN gator.feeds ON inserted_feed_follow.feed_id = gator.feeds.id
INNER JOIN gator.users ON inserted_feed_follow.user_id = gator.users.id;
--


-- name: GetFeedFollowsForUser :many
SELECT 
    gator.feed_follows.feed_id,
    gator.feed_follows.user_id,
    gator.feed_follows.created_at,
    gator.feeds.name AS feed_name, 
    gator.users.name AS user_name
FROM gator.feed_follows
INNER JOIN gator.feeds ON gator.feed_follows.feed_id = gator.feeds.id
INNER JOIN gator.users ON gator.feed_follows.user_id = gator.users.id
WHERE gator.feed_follows.user_id = $1;
--





