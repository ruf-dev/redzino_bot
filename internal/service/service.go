package service

import (
	"context"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type Service interface {
	UserService() UserService
}

type UserService interface {
	InitUser(ctx context.Context, tgId int64) error
	GetBalance(ctx context.Context, tgId int64) (domain.Balance, error)
	AccountRoll(ctx context.Context, bc domain.BalanceChange) (err error)
}
