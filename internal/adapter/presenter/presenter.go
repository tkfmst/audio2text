package presenter

import (
	boundry_output "example.com/audio2text/internal/app/boundary/output"
	"example.com/audio2text/internal/app/data/output"
)

type ResultPresenter struct {
	dest IO
}

func NewResultPresenter(io IO) boundry_output.Audio2Text {
	pr := new(ResultPresenter)
	pr.dest = io

	return pr
}

func (pr *ResultPresenter) OutputStdOut(r output.Result) {
	if r.NonEmpty() {
		pr.dest.Output(r.String())
	}
}
