package db

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type UserProvider struct {
	db sqldb.DB
}

func NewUserProvider(db sqldb.DB) *UserProvider {
	return &UserProvider{
		db: db,
	}
}

func (p *UserProvider) Create(ctx context.Context, data domain.User) (out domain.User, err error) {
	err = p.db.QueryRowContext(ctx, `
		INSERT INTO users 
			   (tg_id, username, balance)
		VALUES (   $1,      $2,       $3)
		ON CONFLICT(tg_id)
		DO UPDATE 
		SET username = excluded.username
		RETURNING tg_id`, data.TgId, data.UserName, data.Balance).
		Scan(&out.TgId)
	if err != nil {
		return out, wrapPgError(err)
	}

	return out, nil
}

func (p *UserProvider) Get(ctx context.Context, tgId int64, opt ...storage.Opt) (user domain.User, err error) {
	q := `
		SELECT
		    tg_id,
		    balance,
		    permission_bit_map,
		    lucky_number
		FROM users 
		WHERE tg_id = $1
		`

	if len(opt) > 0 && opt[0].ForUpdate {
		q += `
		FOR UPDATE`
	}

	err = p.db.QueryRowContext(ctx, q, tgId).
		Scan(
			&user.TgId,
			&user.Balance,
			&user.PermissionsBitMap,
			&user.LuckyNumber,
		)
	if err != nil {
		return user, rerrors.Wrap(err, "error reading user from database")
	}

	return user, nil
}

func (p *UserProvider) ApplyBalanceChange(ctx context.Context, tgId int64, price int) error {
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

func (p *UserProvider) WithTx(tx *sql.Tx) storage.Users {
	return NewUserProvider(tx)
}
