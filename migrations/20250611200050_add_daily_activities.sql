-- +goose Up
-- +goose StatementBegin
CREATE TABLE daily_activities
(
    user_id      INTEGER REFERENCES users(tg_id),
    last_goyda   TIMESTAMP DEFAULT (now() - INTERVAL '1 day') NOT NULL ,
    total_goyda  INTEGER DEFAULT 0
);

INSERT INTO daily_activities (user_id)
SELECT tg_id FROM users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
