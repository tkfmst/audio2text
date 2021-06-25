package adapter

import (
	"log"

	boundry_input "example.com/audio2text/internal/app/boundary/input"
	"example.com/audio2text/internal/app/data/input"
)

type Controller struct {
	usecase boundry_input.Audio2Text
}

func NewController(usecase boundry_input.Audio2Text) Controller {
	c := new(Controller)
	c.usecase = usecase
	return *c
}

func (c *Controller) ConvertAudio2textFromFile(filePath string) {

	file, err := input.NewFile(filePath)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	c.usecase.FromFile(file)
}
