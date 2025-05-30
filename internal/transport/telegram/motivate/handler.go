package motivate

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/media"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/service"
)

const Command = "/motivate"

type Handler struct {
	userService       service.UserService
	motivationService service.MotivationService
}

func New(srv service.Service) *Handler {
	return &Handler{
		userService:       srv.UserService(),
		motivationService: srv.MotivationService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	motivation, err := h.motivationService.GetMotivation(in.Ctx, in.Chat.ID)
	if err != nil {
		return rerrors.Wrap(err, "error when getting motivation")
	}

	if motivation == nil {
		return out.SendMessage(response.NewMessage("Больше видосов нет, хозяин :("))
	}

	videoMessage := response.NewMessage("",
		response.WithMedia(media.Video{
			FileID: motivation.TgFileId,
		}))

	err = out.SendMessage(videoMessage)
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}

func (h *Handler) GetDescription() string {
	return "Наполнится мотивацией"
}

func (h *Handler) GetCommand() string {
	return Command
}
