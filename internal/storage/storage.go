package storage

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

var (
	ErrDuplicated = rerrors.New("duplicated record")
)

type Data interface {
	Users() Users
	Motivations() Motivations
	Chats() Chats

	WithTx(tx *sql.Tx) Data
}

type Users interface {
	Create(ctx context.Context, data domain.User) (domain.User, error)
	Get(ctx context.Context, tgId int64) (domain.User, error)
	ApplyBalanceChange(ctx context.Context, id int64, price int) error

	WithTx(tx *sql.Tx) Users
}

type Motivations interface {
	Save(ctx context.Context, motivation *domain.Motivation) error
	PopForChat(ctx context.Context, id int64) (domain.Motivation, error)
	PushToAllChats(ctx context.Context, motivation domain.Motivation) error

	WithTx(tx *sql.Tx) Motivations
}

type Chats interface {
	Create(ctx context.Context, data domain.Chat) error
	Get(ctx context.Context, tgChatId int64) (chat domain.Chat, err error)
	ToggleMute(ctx context.Context, id int64) error

	WithTx(tx *sql.Tx) Chats
}
