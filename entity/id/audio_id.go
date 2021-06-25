package id

import (
	"fmt"
	"strconv"
	"time"
)

type AudioID struct {
	val string
}

func NewAudioID() *AudioID {
	aid := new(AudioID)
	aid.val = fmt.Sprintf("a%s", strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	return aid

}

func (aid *AudioID) ID() *AudioID {
	return aid
}
