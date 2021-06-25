package output

import "example.com/audio2text/app/data/output"

type Audio2Text interface {
	Write(r output.Result)
}
