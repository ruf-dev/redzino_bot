package defaulthandler

import (
	"go.redsock.ru/toolbox"

	"github.com/ruf-dev/redzino_bot/internal/domain"
)

var (
	fruitsVals = []int{1, 22, 43}
)

const (
	jackpotVal = 64
)

func getSlotPrice(val int) domain.SlotValue {
	if val == jackpotVal {
		return domain.SpinSlotJackpot
	}

	if toolbox.Contains(fruitsVals, val) {
		return domain.SpinSlotFruit
	}

	return domain.SpinSlotUnLuck
}
