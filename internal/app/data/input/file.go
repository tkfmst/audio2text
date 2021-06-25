package input

import (
	"os"

	"github.com/pkg/errors"
)

type File struct {
	file *os.File
}

func NewFile(path string) (*File, error) {
	if f, err := os.Open(path); err != nil {
		return nil, errors.WithStack(err)
	} else {
		fi := new(File)
		fi.file = f

		return fi, nil
	}
}

func (f *File) ToOsFile() *os.File {
	return f.file
}
