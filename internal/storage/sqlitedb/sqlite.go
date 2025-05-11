package sqlitedb

import (
	"database/sql"
)

type Provider struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		db: db,
	}
}
