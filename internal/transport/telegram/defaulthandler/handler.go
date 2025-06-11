package defaulthandler

import (
	"strings"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/media"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/domain/errors"
	"github.com/ruf-dev/redzino_bot/internal/service"
)

var goydaVariants = []string{
	"гойда",
	"goyda",
	"goida",
}

type Handler struct {
	userService       service.UserService
	motivationService service.MotivationService
	dailyActivities   service.DailyActivitiesService

	dicesHandler dicesHandler
}

func New(srv service.Service) *Handler {
	return &Handler{
		userService:       srv.UserService(),
		motivationService: srv.MotivationService(),
		dailyActivities:   srv.DailyActivitiesService(),

		dicesHandler: dicesHandler{
			userService: srv.UserService(),
		},
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) error {
	if in.Dice != nil {
		return h.dicesHandler.Handle(in, out)
	}

	if in.Video != nil {
		return h.handleVideo(in, out)
	}

	if len(in.Text) < 256 && containsGoyda(in.Text) {
		return h.handleGoyda(in, out)
	}

	return nil
}

func (h *Handler) GetDescription() string {
	return "returns current app version as a response"
}

func (h *Handler) handleVideo(in *model.MessageIn, out tgapi.Chat) error {
	mot := domain.Motivation{
		AuthorTgId: in.From.ID,
		TgFileId:   in.Video.FileID,
	}

	var msg *response.MessageOut

	err := h.motivationService.Save(in.Ctx, mot)
	if err == nil {
		msg = response.NewMessage("Видео сохранено")
	} else {
		if !rerrors.Is(err, errors.ErrAccessDenied) {
			return rerrors.Wrap(err)
		}

		msg = response.NewMessage(err.Error())
	}

	if msg != nil {
		msg.ReplyMessageId = int64(in.MessageID)
		return out.SendMessage(msg)
	}

	return nil
}

func (h *Handler) handleGoyda(in *model.MessageIn, out tgapi.Chat) error {
	goydaResp, err := h.dailyActivities.Goyda(in.Ctx, in.From.ID)
	if err != nil {
		return rerrors.Wrap(err)
	}

	var msg tgapi.MessageOut

	if goydaResp.TgFileId != nil {
		msg = response.NewMessage("",
			response.WithMedia(media.Video{
				Caption: "ГОЙДА БРАТЬЯ!",
				FileID:  *goydaResp.TgFileId,
			}))
	} else {
		m := response.NewMessage("Видео нет, но баллы начислены")
		if !goydaResp.ChipsAccounted {
			m.Text = "На сегодня гойды хватит. Крути и приходи завтра"
		}

		m.ReplyMessageId = int64(in.MessageID)

		msg = m
	}

	return out.SendMessage(msg)
}

func containsGoyda(text string) bool {
	text = strings.ToLower(text)

	for _, goyda := range goydaVariants {
		if strings.Contains(text, goyda) {
			return true
		}
	}

	return false
}
