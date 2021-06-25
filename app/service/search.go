package service

import (
	"example.com/audio2text/entity"
	"example.com/audio2text/util/slicestring"
)

const (
	// Idx of the result word
	searchedTargetIdx = 5
)

type Search interface {
	Run(word entity.Word) (*entity.Searched, bool)
	Finalize() []entity.Searched
}

// pos: overall position
type SearchWith5WordsBeforeAndAfter struct {
	result   []entity.Word
	keywords []string
	pos      int
}

func NewSearchWith5WordsBeforeAndAfter(keywords []string) Search {
	s := new(SearchWith5WordsBeforeAndAfter)
	// 最初から11にしておいてshiftしてもエラーにならないように
	s.result = make([]entity.Word, 11)
	s.keywords = keywords
	s.pos = 0

	return s
}

func (s *SearchWith5WordsBeforeAndAfter) Run(word entity.Word) (*entity.Searched, bool) {
	s.pos += 1
	_ = s.popAndPush(word)

	target := s.result[searchedTargetIdx]

	if len(s.result) == 11 && target != nil && slicestring.Contains(s.keywords, target.String()) {
		return entity.NewSearched(s.result, searchedTargetIdx, s.pos-searchedTargetIdx), true
	}
	return nil, false
}

// Finalize searches for SearchWith5WordsBeforeAndAfter.result[6:]
func (s *SearchWith5WordsBeforeAndAfter) Finalize() []entity.Searched {
	result := make([]entity.Searched, searchedTargetIdx)
	for i := 0; i < searchedTargetIdx; i++ {
		r, ok := s.Run(nil)
		if ok {
			result = append(result, *r)
		}
	}
	return result
}

func (s *SearchWith5WordsBeforeAndAfter) popAndPush(word entity.Word) entity.Word {
	w := s.result[0]
	s.result = append(s.result[1:], word)
	return w
}
