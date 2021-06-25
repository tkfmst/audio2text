package api

import (
	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/entity/id"
)

type Converter interface {
	Send(buf []byte) error
	CloseSend() error
	Recv(aid *id.AudioID) (entity.Text, error)
}
