package db

import (
	"context"
	"database/sql"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type MotivationsProvider struct {
	db sqldb.DB
}

func NewMotivationsProvider(db sqldb.DB) *MotivationsProvider {
	return &MotivationsProvider{
		db: db,
	}
}

func (m *MotivationsProvider) Save(ctx context.Context, motivation *domain.Motivation) error {
	err := m.db.QueryRowContext(ctx, `
		INSERT INTO motivations
		    	( tg_file_id, author_tg_id) 
		VALUES 	(         $1,           $2) 
		ON CONFLICT (tg_file_id) DO NOTHING
		RETURNING id
		`,
		motivation.TgFileId, motivation.AuthorTgId).
		Scan(&motivation.Id)
	if err != nil {
		return wrapPgError(err)
	}

	return nil
}

func (m *MotivationsProvider) PopForChat(ctx context.Context, chatId int64) (domain.Motivation, error) {
	out := domain.Motivation{}

	err := m.db.QueryRowContext(ctx, `
	WITH _record AS (
		SELECT
			motivation_id
		FROM motivation_queue
		WHERE tg_chat_id = $1
		AND not is_sent
		ORDER BY random()
		LIMIT 1
	), pop_from_q AS (
		UPDATE motivation_queue mq
		SET is_sent = true
		FROM _record
		WHERE mq.tg_chat_id = $1
		AND   mq.motivation_id = _record.motivation_id
		RETURNING mq.motivation_id
	)
	SELECT
		tg_file_id
	FROM motivations AS m
	INNER JOIN pop_from_q AS poped ON poped.motivation_id  = m.id
	`, chatId).
		Scan(&out.TgFileId)
	if err != nil {
		return domain.Motivation{}, wrapPgError(err)
	}

	return out, nil
}

func (m *MotivationsProvider) PushToAllChats(ctx context.Context, motivation domain.Motivation) error {
	_, err := m.db.ExecContext(ctx, `
		WITH _chats AS (
			SELECT
				tg_chat_id
			FROM chats
			WHERE NOT is_muted
		)
		INSERT INTO motivation_queue
		   	SELECT
			   	c.tg_chat_id,
				$1,
				false
			FROM _chats AS c
		ON CONFLICT (tg_chat_id, motivation_id) DO NOTHING
`, motivation.Id)
	if err != nil {
		return wrapPgError(err)
	}

	return nil
}

func (m *MotivationsProvider) WithTx(tx *sql.Tx) storage.Motivations {
	return NewMotivationsProvider(tx)
}
