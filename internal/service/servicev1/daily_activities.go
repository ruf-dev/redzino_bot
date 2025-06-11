package servicev1

import (
	"context"
	"database/sql"
	"sync/atomic"
	"time"

	errors "go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
	"github.com/ruf-dev/redzino_bot/internal/storage/tx_manager"
)

type DailyActivitiesService struct {
	txManager *tx_manager.TxManager

	dailyActivitiesStorage storage.DailyActivities
	userStorage            storage.Users

	goydaPrize atomic.Int32
}

func NewDailyActivitiesService(txManager *tx_manager.TxManager, data storage.Data) *DailyActivitiesService {
	da := &DailyActivitiesService{
		txManager: txManager,

		dailyActivitiesStorage: data.DailyActivities(),
		userStorage:            data.Users(),
	}

	da.goydaPrize.Store(50)

	return da
}

func (d *DailyActivitiesService) Goyda(ctx context.Context, userId int64) (domain.GoydaResponse, error) {
	resp := domain.GoydaResponse{
		TgFileId:       nil,
		ChipsAccounted: false,
	}

	err := d.txManager.Execute(func(tx *sql.Tx) error {
		dailyActivitiesStorage := d.dailyActivitiesStorage.WithTx(tx)
		userStorage := d.userStorage.WithTx(tx)

		lastGoyda, err := dailyActivitiesStorage.LastGoyda(ctx, userId)
		if err != nil {
			return errors.Wrap(err)
		}

		if lastGoyda.After(time.Now().Add(-24 * time.Hour).UTC()) {
			return nil
		}

		err = userStorage.ApplyBalanceChange(ctx, userId, int(d.goydaPrize.Load()))
		if err != nil {
			return errors.Wrap(err)
		}

		err = dailyActivitiesStorage.AccountGoyda(ctx, userId, time.Now().UTC())
		if err != nil {
			return errors.Wrap(err)
		}

		resp.ChipsAccounted = true
		return nil
	})
	if err != nil {
		return resp, errors.Wrap(err, "error doing goyda")
	}

	if !resp.ChipsAccounted {
		return resp, nil
	}

	return resp, nil
}
