-- +goose Up
ALTER TABLE feed_follows
ADD COLUMN new_feed_id UUID;

ALTER TABLE feed_follows
DROP COLUMN feed_id;

ALTER TABLE feed_follows
RENAME COLUMN new_feed_id TO feed_id;

ALTER TABLE feed_follows
ADD CONSTRAINT fk_feed_id FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE;

-- +goose Down
DROP TABLE feed_follows;