package entity

import (
	"fmt"
	"io"
	"log"
	"os"

	"example.com/audio2text/internal/entity/id"
	"github.com/pkg/errors"
)

type Audio interface {
	ID() *id.AudioID
	Read() ([]byte, error)
	String() string
}

type AudioFile struct {
	*id.AudioID
	file *os.File
	buf  []byte
}

func NewAudioFile(file *os.File) Audio {
	a := new(AudioFile)
	a.file = file
	a.AudioID = id.NewAudioID()
	a.buf = make([]byte, 1024)

	return a
}

func (af *AudioFile) Finalize() {
	if err := af.file.Close(); err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}

func (af *AudioFile) Read() ([]byte, error) {
	n, err := af.file.Read(af.buf)
	if n > 0 {
		return af.buf[:n], nil
	}
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("could not read from %s: %v", af.file.Name(), err))
	}

	return nil, errors.WithStack(fmt.Errorf("unknown file(%s) read error", af.file.Name()))
}

func (af *AudioFile) String() string {
	return fmt.Sprintf("AudioFile(ID=%s, Path=%s)", af.ID(), af.file.Name())
}
