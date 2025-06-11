package domain

import (
	"time"
)

type Chat struct {
	TgId  int64
	Title string

	LastMotivation time.Time
	IsMuted        bool
}

type ScheduledMessage struct {
	Id     int64
	ChatId int64
	Text   string
}

type ScheduledMessageState string

const (
	ScheduledMessageStateWait         ScheduledMessageState = "wait"
	ScheduledMessageStateTaken        ScheduledMessageState = "taken"
	ScheduledMessageStateSent         ScheduledMessageState = "sent"
	ScheduledMessageStateErrorSending ScheduledMessageState = "error_sending"
)
