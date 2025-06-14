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
	dailyActivities   *DailyActivitiesProvider
	settings          *SettingsProvider
}

func NewProvider(db sqldb.DB) *Provider {
	return &Provider{
		users:             NewUserProvider(db),
		motivations:       NewMotivationsProvider(db),
		chats:             NewChatsProvider(db),
		scheduledMessages: NewScheduledMessages(db),
		dailyActivities:   NewDailyActivitiesProvider(db),
		settings:          NewSettingsProvider(db),
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

func (p *Provider) DailyActivities() storage.DailyActivities {
	return p.dailyActivities
}

func (p *Provider) Settings() storage.Settings {
	return p.settings
}

func (p *Provider) WithTx(tx *sql.Tx) storage.Data {
	return NewProvider(tx)
}
