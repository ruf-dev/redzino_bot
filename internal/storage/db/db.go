package db

import (
	"database/sql"

	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type Provider struct {
	users       *UserProvider
	motivations *MotivationsProvider
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		users:       NewUserProvider(db),
		motivations: NewMotivationsProvider(db),
	}
}

func (p *Provider) Users() storage.Users {
	return p.users
}

func (p *Provider) Motivations() storage.Motivations {
	return p.motivations
}
