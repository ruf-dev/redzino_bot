-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS motivations
(
    id           SERIAL PRIMARY KEY,
    tg_file_id   TEXT UNIQUE,
    author_tg_id INTEGER REFERENCES users(tg_id)
);

ALTER TABLE users
    ADD COLUMN permission_bit_map BIGINT DEFAULT 0;

CREATE TABLE IF NOT EXISTS chats
(
    tg_chat_id      BIGINT PRIMARY KEY,
    last_motivation TIMESTAMP,
    is_muted        BOOL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS motivation_queue
(
    tg_chat_id    BIGINT REFERENCES chats (tg_chat_id),
    motivation_id INTEGER REFERENCES motivations (id),
    is_sent       BOOL,
    UNIQUE (tg_chat_id, motivation_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
