package id

import (
	"fmt"
	"strconv"
	"time"
)

type TextID struct {
	val string
}

func NewTextID() *TextID {
	tid := new(TextID)
	tid.val = fmt.Sprintf("t%s", strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	return tid

}

func (tid *TextID) ID() *TextID {
	return tid
}
