package mute

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/service"
)

const Command = "/mute"

type Handler struct {
	chatService service.ChatService
}

func New(srv service.Service) *Handler {
	return &Handler{
		chatService: srv.ChatService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	err := h.chatService.ToggleMute(in.Ctx, in.Chat.ID)
	if err != nil {
		return rerrors.Wrap(err, "error when getting motivation")
	}

	return nil
}

//func (h *Handler) GetDescription() string {
//	return "Включить/выключить оповещения"
//}

func (h *Handler) GetCommand() string {
	return Command
}
