package domain

type Balance struct {
	Total int64
}

type SlotValue int

const (
	SpinSlotUnLuck SlotValue = iota
	SpinSlotFruit
	SpinSlotJackpot
)

type SlotsSpin struct {
	TgId   int64
	Result SlotValue
}

type SlotSpinResult struct {
	IsNotEnoughBalance bool
}

type DiceRoll struct {
	TgId   int64
	Result int
}

type DiceResult int

const (
	DiceRollFailed DiceResult = iota
	DiceRollMatch
)
