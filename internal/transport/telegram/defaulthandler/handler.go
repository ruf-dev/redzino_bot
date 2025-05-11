package defaulthandler

import (
	"runtime"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
)

type Handler struct {
	version string
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(req *model.MessageIn, out tgapi.Chat) error {
	if req.Dice == nil {
		return nil
	}

	runtime.Gosched()
	time.Sleep(1700 * time.Millisecond)

	price := getPrice(req.Dice.Value)
	switch price {
	case jackpotVal:
		return out.SendMessage(response.NewMessage("Грабанул, красавчик!"))
	case fruit:
		return out.SendMessage(response.NewMessage("Лови фруктик"))
	default:
		return out.SendMessage(response.NewMessage("В следующий раз"))
	}

}

func (h *Handler) GetDescription() string {
	return "returns current app version as a response"
}
