package servicev1

import (
	"context"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/domain/errors"
	"github.com/ruf-dev/redzino_bot/internal/domain/permissions"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type MotivationService struct {
	usersStorage       storage.Users
	motivationsStorage storage.Motivations
}

func NewMotivationService(data storage.Data) *MotivationService {
	return &MotivationService{
		usersStorage:       data.Users(),
		motivationsStorage: data.Motivations(),
	}
}

func (m *MotivationService) Save(ctx context.Context, motivation domain.Motivation) (err error) {
	user, err := m.usersStorage.Get(ctx, motivation.AuthorTgId)
	if err != nil {
		return rerrors.Wrap(err, "can't get user from storage")
	}

	if !user.HasPermission(permissions.PermissionAddVideo) {
		return rerrors.Wrap(errors.ErrAccessDenied)
	}

	err = m.motivationsStorage.Save(ctx, motivation)
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}

func (m *MotivationService) GetMotivation(ctx context.Context, chatId int64) (domain.Motivation, error) {

	return domain.Motivation{
		TgFileId: "BAACAgQAAxkBAAMsaDWt4NtSBNdoQFfQ132tcu6jjvIAAtMHAAJmmmxRHmMjaNLtZ0w2BA",
	}, nil
}
