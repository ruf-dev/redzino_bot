package storage

import (
	"context"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

var (
	ErrDuplicated = rerrors.New("duplicated record")
)

type Data interface {
	Users() Users
}

type Users interface {
	Create(ctx context.Context, data domain.User) (domain.User, error)
	Get(ctx context.Context, tgId string) (domain.User, error)
}
