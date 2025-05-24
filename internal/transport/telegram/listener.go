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
	uc := newUserCache()
	{
		s.bot.ExternalContext = func(in *model.MessageIn) context.Context {
			if uc.exists(in.From.ID) {
				return in.Ctx
			}

			err := srv.UserService().InitUser(in.Ctx, in.From.ID)
			if err != nil {
				logrus.Error(rerrors.Wrap(err, "error initializing user"))
				return in.Ctx
			}

			uc.add(in.From.ID)
			return in.Ctx
		}
		// Add handlers here
		s.bot.AddCommandHandler(version.New(cfg.AppInfo.Version))
		s.bot.AddCommandHandler(start.New(srv))
		s.bot.AddCommandHandler(balance.New(srv))

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

type userCache struct {
	users map[int64]struct{}
	sync.RWMutex
}

func newUserCache() *userCache {
	return &userCache{
		users: make(map[int64]struct{}),
	}
}

func (uc *userCache) add(id int64) {
	uc.Lock()
	uc.users[id] = struct{}{}
	uc.Unlock()
}

func (uc *userCache) exists(id int64) bool {
	uc.RLock()
	_, ok := uc.users[id]
	uc.RUnlock()

	return ok
}
