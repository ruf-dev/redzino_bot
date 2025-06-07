package domain

type Balance struct {
	Total int64
}

type SlotSpinResult int

const (
	SpinSlotUnLuck SlotSpinResult = iota
	SpinSlotFruit
	SpinSlotJackpot
)

type SlotsSpin struct {
	TgId   int64
	Result SlotSpinResult
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
