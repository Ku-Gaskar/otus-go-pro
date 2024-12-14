package storage

import (
	"errors"
	"time"
)

var (
	ErrEventAlreadyExists = errors.New("err event already exists")
	ErrEventNotFound      = errors.New("err event not found")
)

type Event struct {
	ID              string
	Title           string
	DateTimeEvent   time.Time
	DateTimeEnd     time.Time
	Description     string
	UserID          string
	PrefixTimeEvent time.Time
}
