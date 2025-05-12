package balance

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

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
	h.userService.GetBalance(in.Ctx, in.From.ID)
	return out.SendMessage(response.NewMessage("Ноль нахуй"))
}

func (h *Handler) GetDescription() string {
	return "Показать баланс"
}

func (h *Handler) GetCommand() string {
	return Command
}
