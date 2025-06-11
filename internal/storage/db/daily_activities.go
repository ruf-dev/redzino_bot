package db

import (
	"context"
	"database/sql"
	"time"

	errors "go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type DailyActivitiesProvider struct {
	db sqldb.DB
}

func (d *DailyActivitiesProvider) LastGoyda(ctx context.Context, userId int64) (time.Time, error) {
	var t time.Time

	err := d.db.QueryRowContext(ctx, `
		SELECT 
			last_goyda
		FROM daily_activities
		WHERE user_id = $1
	`, userId).
		Scan(&t)
	if err != nil {
		return t, errors.Wrap(err, "error getting last goyda from database")
	}

	return t, nil
}

func (d *DailyActivitiesProvider) AccountGoyda(ctx context.Context, userId int64, t time.Time) error {
	_, err := d.db.ExecContext(ctx, `
		UPDATE daily_activities 
		SET last_goyda = $1,
		total_goyda = total_goyda + 1
		WHERE user_id = $2
	`,
		t.UTC(),
		userId)

	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (d *DailyActivitiesProvider) WithTx(tx *sql.Tx) storage.DailyActivities {
	return NewDailyActivitiesProvider(tx)
}

func NewDailyActivitiesProvider(db sqldb.DB) *DailyActivitiesProvider {
	return &DailyActivitiesProvider{
		db: db,
	}
}
