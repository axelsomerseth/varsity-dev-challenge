-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS USER (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    USERNAME TEXT,
    EMAIL TEXT,
    SUBJECT TEXT,
    PICTURE TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE USER;