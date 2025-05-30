package db

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/storage"
)

func wrapPgError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return storage.ErrNotFound
	}

	var pgerr *pq.Error

	if !errors.As(err, &pgerr) {
		return err
	}

	switch pgerr.Code {
	default:
		return rerrors.Wrap(err)
	}
}
