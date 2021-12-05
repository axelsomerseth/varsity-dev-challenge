-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS POST (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    USERID INTEGER,
    POST TEXT,
    CREATED_AT TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE POST;