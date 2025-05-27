package defaulthandler

import (
	"runtime"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/domain/errors"
	"github.com/ruf-dev/redzino_bot/internal/service"
)

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
	if in.Dice != nil {
		return h.handleDice(in, out)
	}

	if in.Video != nil {
		return h.handleVideo(in, out)
	}

	return nil
}

func (h *Handler) GetDescription() string {
	return "returns current app version as a response"
}

func (h *Handler) handleDice(in *model.MessageIn, out tgapi.Chat) error {
	price := getPrice(in.Dice.Value)

	roll := domain.BalanceChange{
		TgId:       in.From.ID,
		RollResult: price,
	}

	err := h.userService.AccountRoll(in.Ctx, roll)
	if err != nil {
		return rerrors.Wrap(err)
	}

	if price == domain.RollPrizeUnLuck {
		return nil
	}

	runtime.Gosched()
	time.Sleep(1700 * time.Millisecond)

	var messageOut *response.MessageOut

	switch price {
	case domain.RollPrizeJackpot:
		messageOut = response.NewMessage("Грабанул, красавчик!")
	case domain.RollPrizeFruit:
		messageOut = response.NewMessage("Лови фруктик")
	}

	if messageOut != nil {
		messageOut.ReplyMessageId = int64(in.MessageID)
		return out.SendMessage(messageOut)
	}

	return nil
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
