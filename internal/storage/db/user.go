package db

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

func (p *UserProvider) Create(ctx context.Context, data domain.User) (out domain.User, err error) {
	err = p.db.QueryRowContext(ctx, `
		INSERT INTO users 
			   (tg_id, balance)
		VALUES (   $1,      $2)
		ON CONFLICT(tg_id)
		DO UPDATE SET tg_id = excluded.tg_id
		RETURNING tg_id`, data.TgId, data.Balance).
		Scan(&out.TgId)
	if err != nil {
		return out, wrapPgError(err)
	}

	return out, nil
}

func (p *UserProvider) Get(ctx context.Context, tgId int64) (user domain.User, err error) {
	err = p.db.QueryRowContext(ctx, `
		SELECT
		    tg_id,
		    balance,
		    permission_bit_map
		FROM users 
		WHERE tg_id = $1`, tgId).
		Scan(
			&user.TgId,
			&user.Balance,
			&user.PermissionsBitMap,
		)
	if err != nil {
		return user, rerrors.Wrap(err, "error reading user from database")
	}

	return user, nil
}

func (p *UserProvider) Inc(ctx context.Context, tgId int64, price int) error {
	return p.updateBalance(ctx, tgId, price)
}

func (p *UserProvider) Decrease(ctx context.Context, tgId int64) error {
	return p.updateBalance(ctx, tgId, -2)
}

func (p *UserProvider) updateBalance(ctx context.Context, tgId int64, balanceChange int) error {
	_, err := p.db.ExecContext(ctx, `
		UPDATE users 
		SET balance = balance + $1
		WHERE tg_id = $2`,
		balanceChange, tgId)
	if err != nil {
		return wrapPgError(err)
	}

	return nil
}
