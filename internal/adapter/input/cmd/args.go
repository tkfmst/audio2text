package cmd

import (
	"flag"
	"fmt"
	"strings"

	"example.com/audio2text/internal/util/logger"
	"github.com/pkg/errors"
)

// Args is command line opiton arguments
type Args struct {
	FilePath   string
	Keywords   []string
	SampleRate int32
}

func NewArgs() (*Args, error) {

	var (
		path       = flag.String("f", "", "audio file path")
		keywords   = flag.String("k", "", "comma-separated search keywords")
		sampleRate = flag.Int("sr", 16000, "audio sample-rate (Hz). default 16000Hz")
	)

	flag.Parse()

	if *path == "" {
		return nil, errors.WithStack(fmt.Errorf("audio file path is empty"))
	}
	if *keywords == "" {
		return nil, errors.WithStack(fmt.Errorf("search keywords is empty"))
	}
	if *sampleRate == -1 {
		return nil, errors.WithStack(fmt.Errorf("sampleRate"))
	}

	args := new(Args)
	args.FilePath = *path
	args.Keywords = strings.Split(*keywords, ",")
	args.SampleRate = int32(*sampleRate)
	logger.Debug("args=", args)

	return args, nil
}
