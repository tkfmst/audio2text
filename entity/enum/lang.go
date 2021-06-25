package enum

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Lang interface {
	String() string
}

// Ja is Lang implementation
type Ja struct {
	code string
}

func NewLang(code string) (Lang, error) {
	switch strings.ToLower(code) {
	case "ja":
		fallthrough
	case "ja-jp":
		ja := new(Ja)
		ja.code = "ja-JP"
		return ja, nil
	default:
		return nil, errors.WithStack(fmt.Errorf("invalid langage code %s", code))
	}
}

func (ja *Ja) String() string {
	return ja.code
}
