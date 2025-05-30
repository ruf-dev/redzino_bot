package domain

import (
	"time"
)

type Chat struct {
	TgId int64

	LastMotivation time.Time
	IsMuted        bool
}
