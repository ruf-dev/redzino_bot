package balance

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
)

const Command = "/balance"

type Handler struct {
	version string
}

func New(version string) *Handler {
	return &Handler{
		version: version,
	}
}

func (h *Handler) GetDescription() string {
	return "Показать баланс"
}

func (h *Handler) GetCommand() string {
	return Command
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	return out.SendMessage(response.NewMessage("Ноль нахуй"))
}
