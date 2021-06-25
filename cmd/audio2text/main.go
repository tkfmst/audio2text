package main

import (
	"log"

	"example.com/audio2text/adapter"
	"example.com/audio2text/adapter/input/cmd"
	"example.com/audio2text/adapter/presenter"
	interactor "example.com/audio2text/app/interactor/gcp"
	"example.com/audio2text/app/service"
	"example.com/audio2text/entity/enum"
	converter_gcp "example.com/audio2text/infra/api/converter/gcp"
	"example.com/audio2text/infra/io"
)

func main() {
	// 仕様:日本語のみ
	lang, err := enum.NewLang("ja-JP")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// 入力
	args, err := cmd.NewArgs()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// usecase生成
	converter, err := converter_gcp.NewLivecaption(lang, args.SampleRate)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	search := service.NewSearchWith5WordsBeforeAndAfter(args.Keywords)
	pr := presenter.NewResultPresenter(io.NewStdOut())

	usecase := interactor.NewAudio2Text(converter, search, pr)

	// controller生成
	c := adapter.NewController(usecase)

	// コマンド実行
	c.ConvertAudio2textFromFile(args.FilePath)
}
