package servicev1

import (
	"context"
	"time"

	"go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type ChatService struct {
	chatStorage storage.Chats
}

func NewChatService(data storage.Data) *ChatService {
	return &ChatService{
		chatStorage: data.Chats(),
	}
}

func (c *ChatService) InitChat(ctx context.Context, chatId int64) error {
	chat := domain.Chat{
		TgId:           chatId,
		LastMotivation: time.Time{},
		IsMuted:        false,
	}

	err := c.chatStorage.Create(ctx, chat)
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}

func (c *ChatService) ToggleMute(ctx context.Context, chatId int64) error {
	err := c.chatStorage.ToggleMute(ctx, chatId)
	if err != nil {
		return rerrors.Wrap(err, "")
	}

	return nil
}

func (c *ChatService) GetChat(ctx context.Context, chatId int64) (*domain.Chat, error) {
	chat, err := c.chatStorage.Get(ctx, chatId)
	if err != nil {
		return nil, rerrors.Wrap(err)
	}

	return &chat, nil
}
