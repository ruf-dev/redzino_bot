package version

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
)

const Command = "/version"

type Handler struct {
	version string
}

func New(version string) *Handler {
	return &Handler{
		version: version,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	return out.SendMessage(response.NewMessage(in.Text + ": " + h.version))
}

func (h *Handler) GetCommand() string {
	return Command
}
