-- +goose Up
-- +goose StatementBegin
ALTER TABLE chats
    ADD COLUMN title TEXT DEFAULT '',
    DROP COLUMN last_motivation;

ALTER TABLE users
    ADD COLUMN username TEXT DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
