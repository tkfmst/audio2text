package service

import (
	"testing"

	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/gomock/mock_entity"
	"github.com/golang/mock/gomock"
)

func TestRun(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	mockWord1 := mock_entity.NewMockWord(c)
	mockWord1.EXPECT().String().Return("ab")

	mockWord2 := mock_entity.NewMockWord(c)
	mockWord2.EXPECT().String().Return("cd")

	mockWord3 := mock_entity.NewMockWord(c)
	mockWord3.EXPECT().String().Return("ef")

	mockWord4 := mock_entity.NewMockWord(c)
	mockWord4.EXPECT().String().Return("gh")

	mockWord5 := mock_entity.NewMockWord(c)
	mockWord5.EXPECT().String().Return("ij")

	mockWord6 := mock_entity.NewMockWord(c)
	mockWord6.EXPECT().String().Return("kl")

	mockWord7 := mock_entity.NewMockWord(c)
	mockWord7.EXPECT().String().Return("mn")

	mockWord8 := mock_entity.NewMockWord(c)
	// mockWord8.EXPECT().String().Return("op")

	mockWord9 := mock_entity.NewMockWord(c)
	// mockWord9.EXPECT().String().Return("qr")

	mockWord10 := mock_entity.NewMockWord(c)
	// mockWord10.EXPECT().String().Return("st")

	mockWord11 := mock_entity.NewMockWord(c)
	// mockWord11.EXPECT().String().Return("vw")

	mockWord12 := mock_entity.NewMockWord(c)
	// mockWord11.EXPECT().String().Return("vw")

	keywords := []string{"ab", "mn", "vw"}
	search := NewSearchWith5WordsBeforeAndAfter(keywords)

	s1, ok1 := search.Run(mockWord1)
	if s1 != nil && !ok1 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s1, ok1)
	}

	s2, ok2 := search.Run(mockWord2)
	if s2 != nil && !ok2 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s2, ok2)
	}

	s3, ok3 := search.Run(mockWord3)
	if s3 != nil && !ok3 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s3, ok3)
	}

	s4, ok4 := search.Run(mockWord4)
	if s4 != nil && !ok4 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s4, ok4)
	}

	s5, ok5 := search.Run(mockWord5)
	if s5 != nil && !ok5 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s5, ok5)
	}

	s6, ok6 := search.Run(mockWord6)
	words6 := []entity.Word{nil, nil, nil, nil, nil, mockWord1, mockWord2, mockWord3, mockWord4, mockWord5, mockWord6}
	searched6 := entity.NewSearched(words6, 5, 6)
	if s6 == searched6 && !ok6 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s6, ok6)
	}

	s7, ok7 := search.Run(mockWord7)
	if s7 != nil && !ok7 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s7, ok7)
	}

	s8, ok8 := search.Run(mockWord8)
	if s8 != nil && !ok8 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s8, ok8)
	}

	s9, ok9 := search.Run(mockWord9)
	if s9 != nil && !ok9 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s9, ok9)
	}

	s10, ok10 := search.Run(mockWord10)
	if s10 != nil && !ok10 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s10, ok10)
	}

	s11, ok11 := search.Run(mockWord11)
	if s11 != nil && !ok11 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s11, ok11)
	}

	s12, ok12 := search.Run(mockWord12)
	words12 := []entity.Word{mockWord2, mockWord3, mockWord4, mockWord5, mockWord6, mockWord7, mockWord8, mockWord9, mockWord10, mockWord11, mockWord12}
	searched12 := entity.NewSearched(words12, 5, 12)
	if s12 == searched12 && !ok12 {
		t.Errorf("invalid search result: searched=%+v, ok=%+v", s12, ok12)
	}
}
