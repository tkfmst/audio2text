package main

import (
	"log"

	"example.com/audio2text/internal/adapter"
	"example.com/audio2text/internal/adapter/input/cmd"
	"example.com/audio2text/internal/adapter/presenter"
	interactor "example.com/audio2text/internal/app/interactor"
	"example.com/audio2text/internal/app/service"
	"example.com/audio2text/internal/entity/enum"
	converter_gcp "example.com/audio2text/internal/infra/api/converter/gcp"
	"example.com/audio2text/internal/infra/io"
)

func main() {
	// 仕様:日本語のみ
	lang, err := enum.NewLang("ja-JP")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// 仕様: 検索単語前後の単語数
	const searchBeforAndAfterWordCount = 5

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
	search := service.NewSearchWithWordsBeforeAndAfter(args.Keywords, searchBeforAndAfterWordCount)
	pr := presenter.NewResultPresenter(io.NewStdOut())

	usecase := interactor.NewAudio2Text(converter, search, pr)

	// controller生成
	c := adapter.NewController(usecase)

	// コマンド実行
	c.ConvertAudio2textFromFile(args.FilePath)
}
