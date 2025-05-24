package servicev1

import (
	"context"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

const (
	defaultInitBalance = 50
	fruitPrice         = 5
	jackpotPrize       = 50
)

type UserService struct {
	userStorage storage.Users
}

func NewUserService(data storage.Data) *UserService {
	return &UserService{
		userStorage: data.Users(),
	}
}

func (u *UserService) InitUser(ctx context.Context, tgId int64) error {
	user := domain.User{
		TgId:    tgId,
		Balance: defaultInitBalance,
	}

	user, err := u.userStorage.Create(ctx, user)
	if err != nil {
		if !rerrors.Is(err, storage.ErrDuplicated) {
			return rerrors.Wrap(err, "error creating user")
		}
	}

	return nil
}

func (u *UserService) GetBalance(ctx context.Context, tgId int64) (domain.Balance, error) {
	user, err := u.userStorage.Get(ctx, tgId)
	if err != nil {
		return domain.Balance{}, rerrors.Wrap(err, "error getting user from storage")
	}

	return domain.Balance{Total: user.Balance}, nil
}

func (u *UserService) AccountRoll(ctx context.Context, bc domain.BalanceChange) (err error) {
	switch bc.RollResult {
	case domain.RollPrizeFruit:
		err = u.userStorage.Inc(ctx, bc.TgId, fruitPrice)
	case domain.RollPrizeJackpot:
		err = u.userStorage.Inc(ctx, bc.TgId, jackpotPrize)
	default:
		err = u.userStorage.Decrease(ctx, bc.TgId)
		if err != nil {
			return rerrors.Wrap(err, "Luck returned back to you. Result wasn't accounted. Problem with database")
		}

	}

	if err != nil {
		return rerrors.Wrap(err, "sorry. Result wasn't accounted")
	}

	return nil
}
