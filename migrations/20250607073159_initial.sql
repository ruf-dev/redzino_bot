-- +goose Up
ALTER TABLE users ADD COLUMN lucky_number INT4 DEFAULT 6;

-- +goose Down
SELECT 'down SQL query';
