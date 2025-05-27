package service

import (
	"context"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type Service interface {
	UserService() UserService
	MotivationService() MotivationService
}

type UserService interface {
	InitUser(ctx context.Context, tgId int64) error
	GetBalance(ctx context.Context, tgId int64) (domain.Balance, error)
	AccountRoll(ctx context.Context, bc domain.BalanceChange) (err error)
}

type MotivationService interface {
	GetMotivation(ctx context.Context, chatId int64) (domain.Motivation, error)
	Save(ctx context.Context, motivation domain.Motivation) (err error)
}
