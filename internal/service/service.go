package service

import (
	"context"
)

type Service interface {
	UserService() UserService
}

type UserService interface {
	InitUser(ctx context.Context, tgId int64) error
}
