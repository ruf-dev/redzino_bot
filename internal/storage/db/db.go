package db

import (
	"database/sql"

	"github.com/ruf-dev/redzino_bot/internal/storage"
)

type Provider struct {
	users *UserProvider
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		users: NewUserProvider(db),
	}
}

func (p *Provider) Users() storage.Users {
	return p.users
}
