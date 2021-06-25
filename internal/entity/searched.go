package entity

import (
	"example.com/audio2text/internal/entity/id"
)

// targetIdx: words[targetIdx].String() == keyword
// pos: overall position of target
type Searched struct {
	*id.SearchedWordsID
	words     []Word
	targetIdx int
	pos       int
}

func NewSearched(words []Word, idx int, pos int) *Searched {
	sw := new(Searched)
	sw.SearchedWordsID = id.NewSearchedWordsID()
	sw.words = words
	sw.targetIdx = idx
	sw.pos = pos

	return sw
}

func (s *Searched) Words() []Word {
	return s.words
}

func (s *Searched) TargetIdx() int {
	return s.targetIdx
}

func (s *Searched) Pos() int {
	return s.pos
}

func (s *Searched) GetStart() string {
	return s.words[s.targetIdx].StartTime().String()
}

func (s *Searched) GetEnd() string {
	return s.words[s.targetIdx].EndTime().String()
}
