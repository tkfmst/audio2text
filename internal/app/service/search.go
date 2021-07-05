package service

import (
	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/util/slicestring"
)

const (
// Idx of the result word
// searchedTargetIdx = 5
)

type Search interface {
	Run(word entity.Word) (*entity.Searched, bool)
	Finalize() []entity.Searched
}

// pos: overall position
type SearchWithWordsBeforeAndAfter struct {
	result                 []entity.Word
	keywords               []string
	pos                    int
	searchedTargetIdx      int
	beforAndAfterWordCount int
}

func NewSearchWithWordsBeforeAndAfter(keywords []string, beforAndAfterWordCount int) Search {

	s := new(SearchWithWordsBeforeAndAfter)
	// 最初から前後単語数+検索ワードにしておいてshiftしてもエラーにならないように
	s.result = make([]entity.Word, beforAndAfterWordCount*2+1)
	s.keywords = keywords
	s.pos = 0
	s.searchedTargetIdx = beforAndAfterWordCount
	s.beforAndAfterWordCount = beforAndAfterWordCount

	return s
}

func (s *SearchWithWordsBeforeAndAfter) Run(word entity.Word) (*entity.Searched, bool) {
	s.pos += 1
	_ = s.popAndPush(word)

	target := s.result[s.searchedTargetIdx]

	if len(s.result) == 11 && target != nil && slicestring.Contains(s.keywords, target.String()) {
		return entity.NewSearched(s.result, s.searchedTargetIdx, s.pos-s.searchedTargetIdx), true
	}
	return nil, false
}

// Finalize searches for SearchWithWordsBeforeAndAfter.result[6:]
func (s *SearchWithWordsBeforeAndAfter) Finalize() []entity.Searched {
	result := make([]entity.Searched, s.searchedTargetIdx)
	for i := 0; i < s.searchedTargetIdx; i++ {
		r, ok := s.Run(nil)
		if ok {
			result = append(result, *r)
		}
	}
	return result
}

func (s *SearchWithWordsBeforeAndAfter) popAndPush(word entity.Word) entity.Word {
	w := s.result[0]
	s.result = append(s.result[1:], word)
	return w
}
