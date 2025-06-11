package db

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	errors "go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/domain"
)

type SettingsProvider struct {
	db sqldb.DB

	m        sync.RWMutex
	settings domain.Settings
}

func NewSettingsProvider(db sqldb.DB) *SettingsProvider {
	sp := &SettingsProvider{
		db: db,
		settings: domain.Settings{
			SlotMachineSettings: domain.SlotMachineSettings{
				RollCost:         5,
				RollFruitPrize:   50,
				RollJackpotPrize: 150,
			},
			DiceSettings: domain.DiceSettings{
				DiceCost: 2,
				DiceWin:  12,
			},
		},
	}

	ctx := context.Background()

	err := sp.fetch(ctx)
	if err != nil {
		logrus.Error(errors.Wrap(err, "error during settings fetch"))
	}

	go func() {
		err := sp.fetch(ctx)
		if err != nil {
			logrus.Error(errors.Wrap(err, "error during settings fetch"))
		}

		ticker := time.NewTicker(5 * time.Second)

		for range ticker.C {
			err = sp.fetch(ctx)
			if err != nil {
				logrus.Error(errors.Wrap(err, "error during settings fetch"))
			}
		}
	}()

	return sp
}

func (d *SettingsProvider) Fetch(ctx context.Context) (domain.Settings, error) {
	err := d.fetch(ctx)
	d.m.RLock()
	defer d.m.RUnlock()

	if err != nil {
		return d.settings, errors.Wrap(err)
	}

	return d.settings, nil
}

func (d *SettingsProvider) Dice() domain.DiceSettings {
	d.m.RLock()
	defer d.m.RUnlock()

	return d.settings.DiceSettings
}

func (d *SettingsProvider) SlotMachine() domain.SlotMachineSettings {
	d.m.RLock()
	defer d.m.RUnlock()

	return d.settings.SlotMachineSettings
}

func (d *SettingsProvider) fetch(ctx context.Context) error {
	d.m.Lock()
	defer d.m.Unlock()

	err := d.db.QueryRowContext(ctx, `
		SELECT 
		    roll_cost,
			roll_fruit_prize,
			roll_jackpot_prize,
			dice_cost,
			dice_win
		FROM settings 
		LIMIT 1`).
		Scan(
			&d.settings.RollCost,
			&d.settings.RollFruitPrize,
			&d.settings.RollJackpotPrize,
			&d.settings.DiceCost,
			&d.settings.DiceWin,
		)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
