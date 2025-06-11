package db

import (
	"context"
	"database/sql"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type ChatsProvider struct {
	db sqldb.DB
}

func NewChatsProvider(db sqldb.DB) *ChatsProvider {
	return &ChatsProvider{
		db: db,
	}
}

func (c *ChatsProvider) Create(ctx context.Context, data domain.Chat) error {
	_, err := c.db.ExecContext(ctx, `
		INSERT INTO chats 
		    	(tg_chat_id, title)  
		VALUES  (        $1,    $2)
		ON CONFLICT(tg_chat_id) DO UPDATE SET title = excluded.title
`, data.TgId, data.Title)
	if err != nil {
		return wrapPgError(err)
	}

	return nil
}

func (c *ChatsProvider) Get(ctx context.Context, tgChatId int64) (chat domain.Chat, err error) {
	chat.TgId = tgChatId

	err = c.db.QueryRowContext(ctx, `
		SELECT 
			last_motivation,
			is_muted
		FROM chats
		WHERE tg_chat_id = $1`).
		Scan(
			&chat.LastMotivation,
			&chat.IsMuted,
		)
	if err != nil {
		return chat, wrapPgError(err)
	}

	return chat, nil
}

func (c *ChatsProvider) ToggleMute(ctx context.Context, id int64) (err error) {
	_, err = c.db.ExecContext(ctx, `
			UPDATE chats
			SET is_muted = not is_muted
			WHERE tg_chat_id = $1`,
		id)

	if err != nil {
		return wrapPgError(err)
	}

	return nil
}

func (c *ChatsProvider) WithTx(tx *sql.Tx) storage.Chats {
	return NewChatsProvider(tx)
}
