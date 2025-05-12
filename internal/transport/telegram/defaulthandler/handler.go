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

type Handler struct {
	userService service.UserService
}

func New(srv service.Service) *Handler {
	return &Handler{
		userService: srv.UserService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	if in.Dice == nil {
		return nil
	}

	runtime.Gosched()
	time.Sleep(1700 * time.Millisecond)

	var messageOut *response.MessageOut

	price := getPrice(in.Dice.Value)
	switch price {
	case domain.RollPrizeJackpot:
		messageOut = response.NewMessage("Грабанул, красавчик!")
	case domain.RollPrizeFruit:
		messageOut = response.NewMessage("Лови фруктик")
	}

	roll := domain.BalanceChange{
		TgId:       in.From.ID,
		RollResult: price,
	}

	err := h.userService.AccountRoll(in.Ctx, roll)
	if err != nil {
		return rerrors.Wrap(err)
	}

	if messageOut != nil {
		return out.SendMessage(messageOut)
	}
	return nil
}

func (h *Handler) GetDescription() string {
	return "returns current app version as a response"
}
