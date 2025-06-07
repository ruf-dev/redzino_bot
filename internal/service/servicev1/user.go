package servicev1

import (
	"context"
	"database/sql"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
	"github.com/ruf-dev/redzino_bot/internal/storage/tx_manager"
)

const (
	defaultInitBalance = 50

	slotFruitPrize   = 25
	slotJackpotPrize = 100
	slotRollPrice    = -2

	diceMatchPrize  = 12
	diceFailedPrice = -2
)

type UserService struct {
	userStorage storage.Users

	txManager *tx_manager.TxManager
}

func NewUserService(data storage.Data, txManager *tx_manager.TxManager) *UserService {
	return &UserService{
		userStorage: data.Users(),
		txManager:   txManager,
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

func (u *UserService) AccountSlotSpin(ctx context.Context, spin domain.SlotsSpin) (err error) {

	var balanceChange int

	switch spin.Result {
	case domain.SpinSlotFruit:
		balanceChange = slotFruitPrize
	case domain.SpinSlotJackpot:
		balanceChange = slotJackpotPrize
	default:
		balanceChange = slotRollPrice
	}

	err = u.userStorage.ApplyBalanceChange(ctx, spin.TgId, balanceChange)
	if err != nil {
		return rerrors.Wrap(err, "sorry. Result wasn't accounted")
	}

	return nil
}

func (u *UserService) AccountDiceRoll(ctx context.Context, roll domain.DiceRoll) (res domain.DiceResult, err error) {
	err = u.txManager.Execute(func(tx *sql.Tx) error {
		userStorage := u.userStorage.WithTx(tx)

		var user domain.User
		user, err = userStorage.Get(ctx, roll.TgId)
		if err != nil {
			return rerrors.Wrap(err, "error getting user info from storage")
		}

		res = diceFailedPrice

		if user.LuckyNumber == roll.Result {
			res = diceMatchPrize
		}

		err = u.userStorage.ApplyBalanceChange(ctx, roll.TgId, int(res))
		if err != nil {
			return rerrors.Wrap(err, "error applying balance")
		}

		return nil
	})

	return res, nil
}
