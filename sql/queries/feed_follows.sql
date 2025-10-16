-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id
;

-- name: GetFeedFollowsForUser :many
SELECT 
    f.name AS feed_name,
    u.name AS user_name
FROM feed_follows
INNER JOIN feeds f ON feed_follows.feed_id = f.id
INNER JOIN users u ON feed_follows.user_id = u.id
WHERE u.name LIKE $1
;