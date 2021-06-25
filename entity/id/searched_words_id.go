package id

import (
	"fmt"
	"strconv"
	"time"
)

type SearchedWordsID struct {
	val string
}

func NewSearchedWordsID() *SearchedWordsID {
	swid := new(SearchedWordsID)
	swid.val = fmt.Sprintf("sw%s", strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	return swid
}

func (swid *SearchedWordsID) ID() *SearchedWordsID {
	return swid
}
