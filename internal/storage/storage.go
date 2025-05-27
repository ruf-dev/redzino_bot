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
	Motivations() Motivations
}

type Users interface {
	Create(ctx context.Context, data domain.User) (domain.User, error)
	Get(ctx context.Context, tgId int64) (domain.User, error)
	Inc(ctx context.Context, id int64, price int) error
	Decrease(ctx context.Context, id int64) error
}

type Motivations interface {
	Save(ctx context.Context, motivation domain.Motivation) error
}
