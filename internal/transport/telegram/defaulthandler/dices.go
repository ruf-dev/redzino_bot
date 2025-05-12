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

func getPrice(val int) domain.RollPrize {
	if val == jackpotVal {
		return domain.RollPrizeJackpot
	}

	if toolbox.Contains(fruitsVals, val) {
		return domain.RollPrizeFruit
	}

	return domain.RollPrizeUnLuck
}
