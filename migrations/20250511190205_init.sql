-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    tg_id   INTEGER PRIMARY KEY,
    balance INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
