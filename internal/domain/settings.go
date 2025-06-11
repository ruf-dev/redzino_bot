package domain

type Settings struct {
	SlotMachineSettings
	DiceSettings
}

type SlotMachineSettings struct {
	RollCost         int
	RollFruitPrize   int
	RollJackpotPrize int
}

type DiceSettings struct {
	DiceCost int
	DiceWin  int
}
