package telegram

import (
	"context"

	client "github.com/Red-Sock/go_tg"

	"github.com/ruf-dev/redzino_bot/internal/config"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/balance"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/defaulthandler"
	"github.com/ruf-dev/redzino_bot/internal/transport/telegram/version"
)

type Server struct {
	bot *client.Bot
}

func NewServer(cfg config.Config, bot *client.Bot) (s *Server) {
	s = &Server{
		bot: bot,
	}

	{
		// Add handlers here
		s.bot.AddCommandHandler(version.New(cfg.AppInfo.Version))
		s.bot.AddCommandHandler(balance.New(cfg.AppInfo.Version))
		s.bot.SetDefaultCommandHandler(defaulthandler.New())
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
