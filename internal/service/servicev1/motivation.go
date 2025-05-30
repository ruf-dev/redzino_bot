package servicev1

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/domain/errors"
	"github.com/ruf-dev/redzino_bot/internal/domain/permissions"
	"github.com/ruf-dev/redzino_bot/internal/storage"
	"github.com/ruf-dev/redzino_bot/internal/storage/tx_manager"
)

type MotivationService struct {
	usersStorage       storage.Users
	motivationsStorage storage.Motivations

	txManager *tx_manager.TxManager
}

func NewMotivationService(data storage.Data, txManager *tx_manager.TxManager) *MotivationService {
	return &MotivationService{
		usersStorage:       data.Users(),
		motivationsStorage: data.Motivations(),

		txManager: txManager,
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

	err = m.txManager.Execute(func(tx *sql.Tx) error {
		motivationsStorage := m.motivationsStorage.WithTx(tx)

		err = motivationsStorage.Save(ctx, &motivation)
		if err != nil {
			return rerrors.Wrap(err)
		}

		err = motivationsStorage.PushToAllChats(ctx, motivation)
		if err != nil {
			return rerrors.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}

func (m *MotivationService) GetMotivation(ctx context.Context, chatId int64) (*domain.Motivation, error) {
	motivation, err := m.motivationsStorage.PopForChat(ctx, chatId)
	if err != nil {
		if !rerrors.Is(err, storage.ErrNotFound) {
			return nil, rerrors.Wrap(err)
		}

		return nil, nil
	}

	return &motivation, nil
}
