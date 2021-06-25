package output

import (
	"fmt"
	"strings"

	"example.com/audio2text/entity"
	"github.com/fatih/color"
)

type Result interface {
	NonEmpty() bool
	String() string
}

type SearchedResult struct {
	val *entity.Searched
}

func NewSearchedResult(initial *entity.Searched) Result {
	sr := new(SearchedResult)
	sr.val = initial
	return sr
}

func (sr *SearchedResult) NonEmpty() bool {
	for _, w := range sr.val.Words() {
		if w != nil && w.String() != "" {
			return true
		}

	}
	return false
}

// TODO Formatterに切り出したい
func (sr *SearchedResult) String() string {
	ss := make([]string, 11)

	for i, w := range sr.val.Words() {
		if w == nil {
			ss[i] = ""
		} else {
			if i == sr.val.TargetIdx() {
				ss[i] = color.RedString(w.String())
			} else {
				ss[i] = w.String()
			}
		}
	}

	ws := strings.Trim(strings.Join(ss, " "), " ")
	if ws == "" {
		return ""
	} else {
		return fmt.Sprintf(
			"[pos: %d, start: %s, end: %s, keyword: %s] %s",
			sr.val.Pos(),
			sr.val.GetStart(),
			sr.val.GetEnd(),
			ss[sr.val.TargetIdx()],
			ws,
		)
	}
}
