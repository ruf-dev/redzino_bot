-- +goose Up
-- +goose StatementBegin
ALTER TABLE scheduled_messages
    DROP CONSTRAINT scheduled_messages_chat_id_fkey;
ALTER TABLE scheduled_messages
    ADD CONSTRAINT scheduled_messages_chat_id_fkey
        FOREIGN KEY (chat_id)
            REFERENCES chats (tg_chat_id)
            ON DELETE CASCADE;

ALTER TABLE scheduled_messages
    DROP CONSTRAINT scheduled_messages_message_id_fkey;
ALTER TABLE scheduled_messages
    ADD CONSTRAINT scheduled_messages_message_id_fkey
        FOREIGN KEY (message_id)
            REFERENCES announcements (id)
            ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
