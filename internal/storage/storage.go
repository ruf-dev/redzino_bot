package storage

import (
	"context"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type Data interface {
}

type Users interface {
	Insert(ctx context.Context, data domain.User) (domain.User, error)
	Get(ctx context.Context, tgId string) (domain.User, error)
}
