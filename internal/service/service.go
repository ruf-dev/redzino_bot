package service

import (
	"context"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type Service interface {
	UserService() UserService
	MotivationService() MotivationService
	ChatService() ChatService
	DailyActivitiesService() DailyActivitiesService
}

type UserService interface {
	InitUser(ctx context.Context, user domain.User) error
	GetBalance(ctx context.Context, tgId int64) (domain.Balance, error)

	AccountSlotSpin(ctx context.Context, bc domain.SlotsSpin) (err error)
	AccountDiceRoll(ctx context.Context, bc domain.DiceRoll) (res domain.DiceResult, err error)
}

type MotivationService interface {
	GetMotivation(ctx context.Context, chatId int64) (*domain.Motivation, error)
	Save(ctx context.Context, motivation domain.Motivation) (err error)
}

type ChatService interface {
	InitChat(ctx context.Context, chat domain.Chat) error
	ToggleMute(ctx context.Context, id int64) error
	GetChat(ctx context.Context, chatId int64) (*domain.Chat, error)
}

type DailyActivitiesService interface {
	Goyda(ctx context.Context, userId int64) (domain.GoydaResponse, error)
}
