package db

import (
	"database/sql"

	"github.com/ruf-dev/redzino_bot/internal/clients/sqldb"
	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type Provider struct {
	users             *UserProvider
	motivations       *MotivationsProvider
	chats             *ChatsProvider
	scheduledMessages *ScheduledMessagesProvider
}

func NewProvider(db sqldb.DB) *Provider {
	return &Provider{
		users:             NewUserProvider(db),
		motivations:       NewMotivationsProvider(db),
		chats:             NewChatsProvider(db),
		scheduledMessages: NewScheduledMessages(db),
	}
}

func (p *Provider) Users() storage.Users {
	return p.users
}

func (p *Provider) Motivations() storage.Motivations {
	return p.motivations
}

func (p *Provider) ScheduledMessages() storage.ScheduledMessages {
	return p.scheduledMessages
}

func (p *Provider) Chats() storage.Chats {
	return p.chats
}

func (p *Provider) WithTx(tx *sql.Tx) storage.Data {
	return NewProvider(tx)
}
