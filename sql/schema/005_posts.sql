-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title text NOT NULL,
    url text NOT NULL UNIQUE,
    description text,
    published_at  TIMESTAMP,
    feed_id UUID not NULL references feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;