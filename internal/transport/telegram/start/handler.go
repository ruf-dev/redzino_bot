package start

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/service"
)

const Command = "/start"

type Handler struct {
	userService service.UserService
}

func New(srv service.Service) *Handler {
	return &Handler{
		userService: srv.UserService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	err := h.userService.InitUser(in.Ctx, in.From.ID)
	if err != nil {
		return rerrors.Wrap(err, "error creating user")
	}

	return out.SendMessage(response.NewMessage("Здарова, заебал 💸💸💸"))
}

func (h *Handler) GetCommand() string {
	return Command
}
