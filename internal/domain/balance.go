package domain

type Balance struct {
	Total int64
}

type RollPrize int

const (
	RollPrizeUnLuck RollPrize = iota
	RollPrizeFruit
	RollPrizeJackpot
)

type BalanceChange struct {
	TgId       int64
	RollResult RollPrize
}
