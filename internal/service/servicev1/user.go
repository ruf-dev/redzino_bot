package servicev1

import (
	"context"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

const defaultInitBalance = 50

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
