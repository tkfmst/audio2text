package entity

import "example.com/audio2text/entity/id"

type Text interface {
	ID() *id.TextID
	AudioID() *id.AudioID
	ReadWord() (Word, bool) // false = EOT
}
