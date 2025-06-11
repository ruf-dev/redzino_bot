package storage

import (
	"context"
	"database/sql"
	"time"

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
	ScheduledMessages() ScheduledMessages
	DailyActivities() DailyActivities

	Settings() Settings

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
	RefreshChatsQueue(ctx context.Context, chatId int64) error

	WithTx(tx *sql.Tx) Motivations
}

type Chats interface {
	Create(ctx context.Context, data domain.Chat) error
	Get(ctx context.Context, tgChatId int64) (chat domain.Chat, err error)
	ToggleMute(ctx context.Context, id int64) error

	WithTx(tx *sql.Tx) Chats
}

type ScheduledMessages interface {
	GetBatchForSending(ctx context.Context, batchSize int) ([]domain.ScheduledMessage, error)
	MarkMessage(ctx context.Context, id int64, state domain.ScheduledMessageState) error

	WithTx(tx *sql.Tx) ScheduledMessages
}

type DailyActivities interface {
	LastGoyda(ctx context.Context, userId int64) (time.Time, error)
	AccountGoyda(ctx context.Context, userId int64, t time.Time) error

	WithTx(tx *sql.Tx) DailyActivities
}

type Settings interface {
	Fetch(ctx context.Context) (domain.Settings, error)

	SlotMachine() domain.SlotMachineSettings
	Dice() domain.DiceSettings
}
