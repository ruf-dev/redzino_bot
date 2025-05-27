package db

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type MotivationsProvider struct {
	db *sql.DB
}

func NewMotivationsProvider(db *sql.DB) *MotivationsProvider {
	return &MotivationsProvider{
		db: db,
	}
}

func (m *MotivationsProvider) Save(ctx context.Context, motivation domain.Motivation) error {
	_, err := m.db.ExecContext(ctx, `
		INSERT INTO motivations
		    	( id, tg_file_id) 
		VALUES 	( $1,         $2) 
		ON CONFLICT (tg_file_id) DO NOTHING`,
		motivation.AuthorTgId, motivation.TgFileId)
	if err != nil {
		return rerrors.Wrap(err, "error saving motivation")
	}

	return nil
}
