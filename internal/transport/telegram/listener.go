package telegram

import (
	"context"
	"sync"

	client "github.com/Red-Sock/go_tg"
	"github.com/Red-Sock/go_tg/model"
	"github.com/sirupsen/logrus"
	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/config"
	"github.com/ruf-dev/redzino_bot/internal/service"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/balance"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/defaulthandler"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/motivate"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/mute"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/start"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/version"
)

type Server struct {
	bot *client.Bot
}

func NewServer(cfg config.Config, bot *client.Bot, srv service.Service) (s *Server) {
	s = &Server{
		bot: bot,
	}

	//TODO RSI-303 replace onto interceptor
	uc := newCache[int64, struct{}]()
	chatCache := newCache[int64, struct{}]()

	{
		s.bot.ExternalContext = func(in *model.MessageIn) context.Context {
			in.Ctx = context.Background()

			if !uc.exists(in.From.ID) {

				err := srv.UserService().InitUser(in.Ctx, in.From.ID)
				if err != nil {
					logrus.Error(rerrors.Wrap(err, "error initializing user"))
					return in.Ctx
				}

				uc.add(in.From.ID, struct{}{})
			}

			if !chatCache.exists(in.Chat.ID) {
				err := srv.ChatService().InitChat(in.Ctx, in.Chat.ID)
				if err != nil {
					logrus.Error(rerrors.Wrap(err, "error initializing chat"))
					return in.Ctx
				}

				chatCache.add(in.Chat.ID, struct{}{})
			}

			return in.Ctx
		}
		// Add handlers here
		s.bot.AddCommandHandler(version.New(cfg.AppInfo.Version))
		s.bot.AddCommandHandler(start.New(srv))
		s.bot.AddCommandHandler(balance.New(srv))
		s.bot.AddCommandHandler(motivate.New(srv))
		s.bot.AddCommandHandler(mute.New(srv))

		s.bot.SetDefaultCommandHandler(defaulthandler.New(srv))
	}

	return s
}

func (s *Server) Start(_ context.Context) error {
	return s.bot.Start()
}

func (s *Server) Stop(_ context.Context) error {
	s.bot.Stop()
	return nil
}

type userCache[T comparable, V any] struct {
	users map[T]V
	sync.RWMutex
}

func newCache[T comparable, V any]() *userCache[T, V] {
	return &userCache[T, V]{
		users: make(map[T]V),
	}
}

func (uc *userCache[T, V]) add(k T, v V) {
	uc.Lock()
	uc.users[k] = v
	uc.Unlock()
}

func (uc *userCache[T, V]) exists(key T) bool {
	uc.RLock()
	_, ok := uc.users[key]
	uc.RUnlock()

	return ok
}
