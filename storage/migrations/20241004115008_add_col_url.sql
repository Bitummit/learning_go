-- +goose Up
-- +goose StatementBegin
ALTER TABLE url
    ADD COLUMN test VARCHAR(256) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE url
    DROP COLUMN test;
-- +goose StatementEnd