package start

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

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
	return out.SendMessage(response.NewMessage("Привет! Это бот Луданхамон Первый! " +
		"Тут можно заработать фишек. Присылай смайлик игрового автомата - 🎰 или костей - 🎲 и выигрывай "))
}

func (h *Handler) GetCommand() string {
	return Command
}
