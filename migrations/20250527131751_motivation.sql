-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS motivations
(
    id         SERIAL,
    tg_file_id TEXT UNIQUE 
);

ALTER TABLE users
    ADD COLUMN permission_bit_map BIGINT DEFAULT 0;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
