package enum

import (
	"testing"
)

func TestLang1(t *testing.T) {
	code1 := "ja"
	lang1, err1 := NewLang(code1)
	if lang1.String() != "ja-JP" || err1 != nil {
		t.Errorf("NewLang should return enum.Ja when the collect lang code is entered: code=%s, err=%+v", lang1, err1)
	}

	code2 := "JA"
	lang2, err2 := NewLang(code2)
	if lang2.String() != "ja-JP" || err2 != nil {
		t.Errorf("NewLang should return enum.Ja when the collect lang code is entered: code=%s, err=%+v", lang2, err2)
	}

	code3 := "ja-jp"
	lang3, err3 := NewLang(code3)
	if lang3.String() != "ja-JP" || err3 != nil {
		t.Errorf("NewLang should return enum.Ja when the collect lang code is entered: code=%s, err=%+v", lang3, err3)
	}
}

func TestLang2(t *testing.T) {
	code := "us"
	_, err := NewLang("us")
	if err == nil {
		t.Errorf("NewLang should return error when invalid input code=%s, err=%+v", code, err)
	}
}
