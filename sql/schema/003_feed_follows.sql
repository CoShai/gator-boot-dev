-- +goose Up
CREATE TABLE feed_follows  (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL references users(id) on delete cascade,
    feed_id UUID not NULL references feeds(id) on delete cascade,
    UNIQUE(user_id,feed_id)
);

-- +goose Down
DROP TABLE feed_follows;