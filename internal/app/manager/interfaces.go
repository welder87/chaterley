package manager

import "chaterley/internal/app/message"

type Session interface {
	SendMessage(msg message.MessageSnapshot) error
	GetID() string
	Close() error
}
