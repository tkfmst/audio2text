package output

import (
	"testing"
	"time"

	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/gomock/mock_entity"
	"github.com/golang/mock/gomock"
)

// TestString1 : Words contains empty entity.Word
func TestString1(t *testing.T) {
	// Words is empty
	emptyWords := make([]entity.Word, 1)
	s := entity.NewSearched(emptyWords, 0, 0)
	sr := NewSearchedResult(s)

	obtained := sr.String()
	expected := ""
	if obtained != expected {
		t.Errorf("Words.String() shoud return empty string when Words.ws is empty slice: obtained=`%+v`, expected=`%+v`", obtained, expected)
	}
}

// Words contains 1 entity.Word
func TestString2(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	d1s, _ := time.ParseDuration("1s")
	d2s, _ := time.ParseDuration("2s")
	mockWord := mock_entity.NewMockWord(c)
	mockWord.EXPECT().String().Return("a")
	mockWord.EXPECT().StartTime().Return(d1s)
	mockWord.EXPECT().EndTime().Return(d2s)

	var ws = []entity.Word{mockWord}
	s := entity.NewSearched(ws, 0, 1)
	sr := NewSearchedResult(s)

	obtained := sr.String()
	expected := "[pos: 1, start: 1s, end: 2s, keyword: a] a"
	if obtained != expected {
		t.Errorf("Words.String() shoud return entity.Word.String() when Words.ws contains 1 entity.Word slice: obtained=`%+v`, expected=`%+v`", obtained, expected)
	}
}

// Words contains multiple entity.Word
func TestString3(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	d1s, _ := time.ParseDuration("1s")
	d2s, _ := time.ParseDuration("2s")

	mockWord1 := mock_entity.NewMockWord(c)
	mockWord1.EXPECT().String().Return("a")

	mockWord2 := mock_entity.NewMockWord(c)
	mockWord2.EXPECT().String().Return("b")
	mockWord2.EXPECT().StartTime().Return(d1s)
	mockWord2.EXPECT().EndTime().Return(d2s)

	var ws = []entity.Word{mockWord1, mockWord2}
	s := entity.NewSearched(ws, 1, 2)
	sr := NewSearchedResult(s)

	obtained := sr.String()
	expected := "[pos: 2, start: 1s, end: 2s, keyword: b] a b"
	if obtained != expected {
		t.Errorf("Words.String() shoud return entity.Word.String() when Words.ws contains 1 entity.Word slice: obtained=`%+v`, expected=`%+v`", obtained, expected)
	}
}

// Words contains nil
func TestString4(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	d1s, _ := time.ParseDuration("1s")
	d2s, _ := time.ParseDuration("2s")

	mockWord1 := mock_entity.NewMockWord(c)
	mockWord1.EXPECT().String().Return("a")
	mockWord1.EXPECT().StartTime().Return(d1s)
	mockWord1.EXPECT().EndTime().Return(d2s)

	mockWord2 := mock_entity.NewMockWord(c)
	mockWord2.EXPECT().String().Return("b")

	var ws = []entity.Word{nil, mockWord1, mockWord2, nil}
	s := entity.NewSearched(ws, 1, 4)
	sr := NewSearchedResult(s)

	obtained := sr.String()
	expected := "[pos: 4, start: 1s, end: 2s, keyword: a] a b"
	if obtained != expected {
		t.Errorf("Words.String() shoud ignore nil: obtained=`%+v`, expected=`%+v`", obtained, expected)
	}
}
