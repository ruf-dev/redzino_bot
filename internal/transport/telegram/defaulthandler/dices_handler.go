package defaulthandler

import (
	"runtime"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/service"
)

type dicesHandler struct {
	userService service.UserService
}

func (h *dicesHandler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	if in.Dice == nil {
		return nil
	}
	switch in.Dice.Emoji {
	case "🎰":
		return h.handleSlots(in, out)
	case "🎲":
		return h.handleDice(in, out)
	}

	return nil
}

func (h *dicesHandler) handleSlots(in *model.MessageIn, out tgapi.Chat) error {

	spin := domain.SlotsSpin{
		TgId:   in.From.ID,
		Result: getSlotPrice(in.Dice.Value),
	}

	err := h.userService.AccountSlotSpin(in.Ctx, spin)
	if err != nil {
		return rerrors.Wrap(err)
	}

	var messageOut *response.MessageOut

	switch spin.Result {
	case domain.SpinSlotJackpot:
		messageOut = response.NewMessage("Грабанул, красавчик!")
	case domain.SpinSlotFruit:
		messageOut = response.NewMessage("Лови фруктик")
	default:
		return nil
	}

	runtime.Gosched()
	time.Sleep(1700 * time.Millisecond)

	if messageOut != nil {
		messageOut.ReplyMessageId = int64(in.MessageID)
		return out.SendMessage(messageOut)
	}

	return nil
}

func (h *dicesHandler) handleDice(in *model.MessageIn, out tgapi.Chat) error {
	roll := domain.DiceRoll{
		TgId:   in.From.ID,
		Result: in.Dice.Value,
	}

	diceRes, err := h.userService.AccountDiceRoll(in.Ctx, roll)
	if err != nil {
		return rerrors.Wrap(err)
	}

	if diceRes == domain.DiceRollMatch {
		return out.SendMessage(response.NewMessage("Совпадение!"))
	}

	return nil

}
