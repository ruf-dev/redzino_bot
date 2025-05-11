package db

import (
	"errors"

	"github.com/lib/pq"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/storage"
)

const (
	errCodeDuplicate = "23502"
)

func wrapPgError(err error) error {
	var pgerr *pq.Error

	if !errors.As(err, &pgerr) {
		return err
	}

	switch pgerr.Code {
	case errCodeDuplicate:
		return rerrors.Wrap(storage.ErrDuplicated, err.Error())
	}

	return err
}
