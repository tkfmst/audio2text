package entity

import (
	"time"

	"example.com/audio2text/internal/entity/id"
)

type Word interface {
	ID() *id.WordID
	TextID() *id.TextID
	StartTime() time.Duration
	EndTime() time.Duration
	String() string
}
