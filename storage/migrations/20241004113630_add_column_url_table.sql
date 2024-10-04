-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-   ADD COLUMN test INTEGER 
SELECT 'down SQL query';
-- +goose StatementEnd
