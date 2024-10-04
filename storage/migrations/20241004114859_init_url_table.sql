-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS url (
        id SERIAL PRIMARY KEY,
        url VARCHAR(256) NOT NULL,
        alias VARCHAR(256) NOT NULL UNIQUE
    );
CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS url;
-- +goose StatementEnd