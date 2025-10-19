-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT f.*, u.name as user_name
  FROM feeds f
  JOIN users u ON f.user_id = u.id
;

-- name: GetFeed :one
SELECT * FROM feeds WHERE url LIKE $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
  SET updated_at = NOW(), last_fetched_at = NOW()
  WHERE id = $1
;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
  ORDER BY last_fetched_at ASC NULLS FIRST
  LIMIT 1
;
