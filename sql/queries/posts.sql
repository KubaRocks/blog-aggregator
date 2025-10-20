-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *
;

-- name: GetPostsForUser :many
SELECT posts.*
    FROM posts
    JOIN feed_follows ff ON ff.feed_id = posts.feed_id
    WHERE ff.user_id = $1
    LIMIT $2
;
