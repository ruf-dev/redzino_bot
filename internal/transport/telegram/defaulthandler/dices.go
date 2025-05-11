package defaulthandler

import (
	"go.redsock.ru/toolbox"
)

var (
	fruitsVals = []int{1, 22, 43}
)

const (
	jackpotVal = 64

	loss = iota
	fruit
	jackpot
)

func getPrice(val int) int {
	if val == jackpotVal {
		return jackpot
	}

	if toolbox.Contains(fruitsVals, val) {
		return fruit
	}

	return loss
}
