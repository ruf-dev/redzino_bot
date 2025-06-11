package cron

import (
	"context"
	"time"

	"github.com/Red-Sock/go_tg"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/sirupsen/logrus"
	errors "go.redsock.ru/rerrors"

	"github.com/ruf-dev/redzino_bot/internal/domain"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

const scheduledMessagesBatchSize = 1

type MessageSenderCron struct {
	messagesProvider storage.ScheduledMessages
	chatsProvider    storage.Chats

	tgApi *go_tg.Bot
}

func NewMessageSender(data storage.Data, tgApi *go_tg.Bot) MessageSenderCron {
	return MessageSenderCron{
		messagesProvider: data.ScheduledMessages(),
		chatsProvider:    data.Chats(),

		tgApi: tgApi,
	}
}

func (m *MessageSenderCron) Start(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)

	err := m.do(ctx)
	if err != nil {
		logrus.Error(errors.Wrap(err))
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := m.do(ctx)
			if err != nil {
				logrus.Error(errors.Wrap(err))
			}
		}
	}
}

func (m *MessageSenderCron) do(ctx context.Context) error {
	batch, err := m.messagesProvider.GetBatchForSending(ctx, scheduledMessagesBatchSize)
	if err != nil {
		return errors.Wrap(err, "error getting batch for update")
	}

	for _, msg := range batch {
		tgMsg := response.NewMessage(msg.Text)
		tgMsg.ChatId = msg.ChatId

		state := domain.ScheduledMessageStateSent
		err = m.tgApi.Send(tgMsg)
		if err != nil {
			logrus.Errorf("could not send message to telegram chat %d: Text: %s",
				tgMsg.ChatId, tgMsg.Text)

			state = domain.ScheduledMessageStateErrorSending
		}

		err = m.messagesProvider.MarkMessage(ctx, msg.Id, state)
		if err != nil {
			logrus.Error(errors.Wrap(err, "error saving message to storage"))
		}
	}

	return nil
}
