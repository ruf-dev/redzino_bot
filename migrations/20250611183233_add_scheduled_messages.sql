-- +goose Up
-- +goose StatementBegin

CREATE TABLE announcements
(
    id      SERIAL PRIMARY KEY,
    message TEXT
);

CREATE TYPE scheduled_message_state AS ENUM ('wait', 'taken', 'sent', 'error_sending', 'muted');

CREATE TABLE scheduled_messages
(
    id         SERIAL PRIMARY KEY,
    chat_id    BIGINT REFERENCES chats (tg_chat_id),
    message_id INT REFERENCES announcements (id),
    state      SCHEDULED_MESSAGE_STATE DEFAULT 'wait'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
