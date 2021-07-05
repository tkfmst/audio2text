package output

import "example.com/audio2text/internal/app/data/output"

type Audio2Text interface {
	OutputStdOut(r output.Result)
}
