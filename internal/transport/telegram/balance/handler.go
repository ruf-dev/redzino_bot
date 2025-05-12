package balance

import (
	"fmt"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/service"
)

const Command = "/balance"

type Handler struct {
	userService service.UserService
}

func New(srv service.Service) *Handler {
	return &Handler{
		userService: srv.UserService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	balance, err := h.userService.GetBalance(in.Ctx, in.From.ID)
	if err != nil {
		return rerrors.Wrap(err)
	}

	return out.SendMessage(response.NewMessage(fmt.Sprintf("Общий баланс: %d фишек", balance.Total)))
}

func (h *Handler) GetDescription() string {
	return "Показать баланс"
}

func (h *Handler) GetCommand() string {
	return Command
}
