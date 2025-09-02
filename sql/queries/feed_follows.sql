-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (
    insert into feed_follows (id,created_at,updated_at,user_id,feed_id)
    VALUES ($1,$2,$3,$4,$5)
    RETURNING *
) 
SELECT inserted_feed_follow.*,feeds.name as feed_name, users.name as user_name
from inserted_feed_follow
join feeds on feeds.id=inserted_feed_follow.feed_id
join users on users.id =inserted_feed_follow.user_id;
--

-- name: GetFeedFollowsForUser :many
select feed_follows.*,users.name as user_name,feeds.name as feed_name from feed_follows
join feeds on feeds.id=feed_follows.feed_id
join users on feed_follows.user_id=users.id
where feed_follows.user_id=$1;
--

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows;
--

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
where user_id=$1 and feed_id=$2;
--

