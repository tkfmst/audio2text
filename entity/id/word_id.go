package id

import (
	"fmt"
	"strconv"
	"time"
)

type WordID struct {
	val string
}

func NewWordID() *WordID {
	wid := new(WordID)
	wid.val = fmt.Sprintf("w%s", strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	return wid
}

func (wid *WordID) ID() *WordID {
	return wid
}
