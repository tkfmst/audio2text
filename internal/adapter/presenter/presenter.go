package presenter

import "example.com/audio2text/internal/app/data/output"

type Presenter interface {
	Write(r output.Result)
}

type ResultPresenter struct {
	dest IO
}

func NewResultPresenter(io IO) Presenter {
	pr := new(ResultPresenter)
	pr.dest = io

	return pr
}

func (pr *ResultPresenter) Write(r output.Result) {
	if r.NonEmpty() {
		pr.dest.Output(r.String())
	}
}
