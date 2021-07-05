package gcp

import (
	"io"
	"log"

	boundry_input "example.com/audio2text/internal/app/boundary/input"
	boundry_output "example.com/audio2text/internal/app/boundary/output"
	"example.com/audio2text/internal/app/data/input"
	"example.com/audio2text/internal/app/data/output"
	"example.com/audio2text/internal/app/if/api"
	"example.com/audio2text/internal/app/service"
	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/util/logger"
)

type Audio2Text struct {
	converter api.Converter
	search    service.Search
	presenter boundry_output.Audio2Text
}

func NewAudio2Text(
	conv api.Converter,
	search service.Search,
	presenter boundry_output.Audio2Text,
) boundry_input.Audio2Text {
	a2t := new(Audio2Text)
	a2t.converter = conv
	a2t.search = search
	a2t.presenter = presenter

	return a2t
}

func (a2t Audio2Text) FromFile(f *input.File) {
	audio := entity.NewAudioFile(f.ToOsFile())

	// Read file & Request api
	go func() {
		for {
			buf, err := audio.Read()
			if err == io.EOF {
				if err := a2t.converter.CloseSend(); err != nil {
					log.Fatalf("%+v", err)
				}
				logger.Debug("Audio.Read finished")
				break
			}
			if err != nil {
				log.Printf("%+v", err)
				continue
			}

			if err := a2t.converter.Send(buf); err != nil {
				log.Fatalf("%+v", err)
			}
		}
	}()

	// Response api & Search keyword
	i := 1
	for {
		i = i + 1
		text, err := a2t.converter.Recv(audio.ID())
		if err == io.EOF {
			logger.Debug("Converter.Recv finished")
			break
		}
		if err != nil {
			log.Fatalf("%+v", err)
		}

		for {
			word, ok := text.ReadWord()
			logger.Debug(word, ok)
			if !ok { // text EOT
				for _, searched := range a2t.search.Finalize() {
					a2t.presenter.OutputStdOut(output.NewSearchedResult(&searched))
				}
				logger.Debug("Search.Run&Finalize finished")
				break
			}

			searched, ok := a2t.search.Run(word)
			logger.Debug("searched:", searched, ok)
			if ok {
				a2t.presenter.OutputStdOut(output.NewSearchedResult(searched))
			}
		}
	}
}
