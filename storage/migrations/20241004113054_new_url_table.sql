-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
    CREATE TABLE IF NOT EXISTS url (
        id SERIAL PRIMARY KEY,
        url VARCHAR(256) NOT NULL,
        alias VARCHAR(256) NOT NULL UNIQUE
    );
    CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS url;
