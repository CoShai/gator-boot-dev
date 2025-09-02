-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name,url,user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
--

-- name: GetFeed :one
SELECT * from feeds
WHERE id = $1 LIMIT 1;
--

-- name: GetFeedByName :one
SELECT * from feeds
WHERE name = $1 LIMIT 1;
--

-- name: GetFeedByUrl :one
SELECT * from feeds
WHERE url = $1 LIMIT 1;
--

-- name: DeleteFeeds :exec
DELETE FROM feeds;
--

-- name: GetFeeds :many
SELECT * from feeds;
--

-- name: GetUserByFeed :one
SELECT users.* from feeds
join users on feeds.user_id =users.id
WHERE feeds.id = $1 LIMIT 1;
--

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id=$1;
--

-- name: GetNextFeedToFetch :one
SELECT * from feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
 