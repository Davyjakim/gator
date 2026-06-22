-- +goose up
CREATE TABLE feeds_follow (
    id UUID PRIMARY KEY,
    user_id UUID,
    feed_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE(user_id, feed_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose down
DROP TABLE feeds_follow;