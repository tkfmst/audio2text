package io

import "fmt"

type StdOut struct {
}

func NewStdOut() *StdOut {
	return new(StdOut)
}

func (so *StdOut) Output(s string) {
	fmt.Println(s)
}
