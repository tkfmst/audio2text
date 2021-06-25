package mock_presenter

import (
	"example.com/audio2text/adapter/presenter"
)

type MockIO struct {
	ch chan string
}

func NewMockIO(ch chan string) presenter.IO {
	io := new(MockIO)
	io.ch = ch
	return io
}

func (mi *MockIO) Output(s string) {
	mi.ch <- s
}
