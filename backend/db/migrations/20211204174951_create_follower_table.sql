-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS FOLLOWER (
    USER_ID INTEGER,
    FOLLOWER_ID INTEGER
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE FOLLOWER;