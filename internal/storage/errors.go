package storage

import (
	"go.redsock.ru/rerrors"
)

var (
	ErrNotFound = rerrors.NewUserError("Not found")
)
