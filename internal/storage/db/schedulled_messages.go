package db

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type ScheduledMessagesProvider struct {
	db sqldb.DB
}

func NewScheduledMessages(db sqldb.DB) *ScheduledMessagesProvider {
	return &ScheduledMessagesProvider{
		db: db,
	}
}

func (s *ScheduledMessagesProvider) GetBatchForSending(ctx context.Context, batchSize int) (
	[]domain.ScheduledMessage, error) {

	rows, err := s.db.QueryContext(ctx, `
		WITH _msg AS (
			SELECT
				id,
				chat_id
			FROM scheduled_messages msg
			INNER JOIN chats ON chats.tg_chat_id = msg.chat_id
			WHERE state = 'wait'
			FETCH FIRST $1 ROWS ONLY
		), taken_msgs as (
		    UPDATE scheduled_messages 
			SET state = 'taken' 
			WHERE id IN (
				SELECT
				    _msg.id
				FROM _msg 
				INNER JOIN chats c ON c.tg_chat_id = _msg.chat_id AND NOT c.is_muted
			)
			RETURNING id
		), muted AS (
		    UPDATE scheduled_messages 
		    SET state = 'muted'
		    WHERE id IN (
				SELECT
				    _msg.id
				FROM _msg 
				INNER JOIN chats c ON c.tg_chat_id = _msg.chat_id AND c.is_muted
			)
		)
		SELECT 
			sm.id, 
			sm.chat_id,
			an.message
		FROM scheduled_messages sm
		INNER JOIN announcements AS an
		ON sm.message_id = an.id
		WHERE sm.id IN (SELECT id FROM taken_msgs)
		`, batchSize)
	if err != nil {
		return nil, rerrors.Wrap(err)
	}
	defer rows.Close()

	out := make([]domain.ScheduledMessage, 0, batchSize)

	for rows.Next() {
		message := domain.ScheduledMessage{}

		err = rows.Scan(
			&message.Id,
			&message.ChatId,
			&message.Text,
		)
		if err != nil {
			return nil, rerrors.Wrap(err, "error scanning row")
		}

		out = append(out, message)
	}

	return out, nil
}

func (s *ScheduledMessagesProvider) MarkMessage(ctx context.Context, id int64, state domain.ScheduledMessageState) error {
	_, err := s.db.ExecContext(ctx, `
		UPDATE scheduled_messages 
		SET state = $1 
		WHERE id = $2
`, state, id)
	if err != nil {
		return rerrors.Wrap(err, "error deleting scheduled messages from database")
	}

	return nil
}

func (s *ScheduledMessagesProvider) WithTx(tx *sql.Tx) storage.ScheduledMessages {
	return NewScheduledMessages(tx)
}
