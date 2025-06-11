package domain

import (
	"github.com/ruf-dev/redzino_bot/internal/domain/permissions"
)

type User struct {
	TgId     int64
	UserName string

	Balance           int64
	PermissionsBitMap int64
	LuckyNumber       int
}

func (u User) HasPermission(perm permissions.Permission) bool {
	return u.PermissionsBitMap&int64(perm) != 0
}
