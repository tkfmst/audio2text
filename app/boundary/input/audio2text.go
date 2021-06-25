package input

import "example.com/audio2text/app/data/input"

type Audio2Text interface {
	FromFile(f *input.File)
}
