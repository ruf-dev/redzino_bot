package db

import (
	"errors"

	"github.com/lib/pq"
	"go.redsock.ru/rerrors"
)

func wrapPgError(err error) error {
	var pgerr *pq.Error

	if !errors.As(err, &pgerr) {
		return err
	}

	switch pgerr.Code {
	default:
		return rerrors.Wrap(err)
	}
}
