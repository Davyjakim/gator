-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feeds_follow
    (id, created_at, updated_at,user_id, feed_id) 
    VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, users.name AS user_name FROM feeds_follow
INNER JOIN users ON users.id = feeds_follow.user_id
INNER JOIN feeds ON feeds.id = feeds_follow.feed_id
WHERE users.id = $1;