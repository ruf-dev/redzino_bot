package sqlitedb

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type UserProvider struct {
	db *sql.DB
}

func NewUserProvider(db *sql.DB) *UserProvider {
	return &UserProvider{
		db: db,
	}
}

func (p *UserProvider) Insert(ctx context.Context, data domain.User) (out domain.User, err error) {
	err = p.db.QueryRowContext(ctx, `
		INSERT INTO users 
			   (tg_id)
		VALUES (   $1)
		RETURNING tg_id`, data.TgId).
		Scan(&out.TgId)
	if err != nil {
		return out, rerrors.Wrap(err, "error upserting user")
	}

	return out, nil
}

func (p *UserProvider) Get(ctx context.Context, tgId string) (user domain.User, err error) {
	err = p.db.QueryRow(`
		SELECT
		    tg_id,
		    balance
		FROM users 
		WHERE tg_id = $1`, tgId).
		Scan(
			&user.TgId,
			&user.Balance,
		)
	if err != nil {
		return user, rerrors.Wrap(err, "error reading user from database")
	}

	return user, nil
}
