package domain

import (
	"github.com/ruf-dev/redzino_bot/internal/domain/permissions"
)

type User struct {
	TgId              int64
	Balance           int64
	PermissionsBitMap int64
}

func (u User) HasPermission(perm permissions.Permission) bool {
	return u.PermissionsBitMap&int64(perm) != 0
}
