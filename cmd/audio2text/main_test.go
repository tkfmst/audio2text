package main

import (
	"io"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"example.com/audio2text/internal/adapter"
	"example.com/audio2text/internal/adapter/presenter"
	interactor "example.com/audio2text/internal/app/interactor"
	"example.com/audio2text/internal/app/service"
	"example.com/audio2text/internal/gomock/mock_api"
	"example.com/audio2text/internal/gomock/mock_entity"
	"example.com/audio2text/internal/mock/mock_presenter"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestMain(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	mockWord1 := mock_entity.NewMockWord(c)
	mockWord2 := mock_entity.NewMockWord(c)
	mockWord3 := mock_entity.NewMockWord(c)
	mockWord4 := mock_entity.NewMockWord(c)
	mockWord5 := mock_entity.NewMockWord(c)
	mockWord6 := mock_entity.NewMockWord(c)
	mockWord7 := mock_entity.NewMockWord(c)
	mockWord8 := mock_entity.NewMockWord(c)
	mockWord9 := mock_entity.NewMockWord(c)
	mockWord10 := mock_entity.NewMockWord(c)
	mockWord11 := mock_entity.NewMockWord(c)
	mockWord12 := mock_entity.NewMockWord(c)
	mockWordEOT := mock_entity.NewMockWord(c)

	d1s, _ := time.ParseDuration("1s")
	d2s, _ := time.ParseDuration("2s")
	d3s, _ := time.ParseDuration("3s")
	d4s, _ := time.ParseDuration("4s")
	d5s, _ := time.ParseDuration("5s")
	d6s, _ := time.ParseDuration("6s")

	mockWord1.EXPECT().String().Return("1abい").AnyTimes()
	mockWord1.EXPECT().StartTime().Return(d1s).AnyTimes()
	mockWord1.EXPECT().EndTime().Return(d2s).AnyTimes()

	mockWord1.EXPECT().String().Return("1abい").AnyTimes()
	mockWord2.EXPECT().String().Return("2cdろ").AnyTimes()
	mockWord3.EXPECT().String().Return("3efは").AnyTimes()
	mockWord4.EXPECT().String().Return("4ghに").AnyTimes()
	mockWord5.EXPECT().String().Return("5ijほ").AnyTimes()

	mockWord6.EXPECT().String().Return("6klへ").AnyTimes()
	mockWord6.EXPECT().StartTime().Return(d3s).AnyTimes()
	mockWord6.EXPECT().EndTime().Return(d4s).AnyTimes()

	mockWord7.EXPECT().String().Return("7mnと").AnyTimes()
	mockWord8.EXPECT().String().Return("8opち").AnyTimes()
	mockWord9.EXPECT().String().Return("9qrり").AnyTimes()
	mockWord10.EXPECT().String().Return("10stぬ").AnyTimes()
	mockWord11.EXPECT().String().Return("11uvる").AnyTimes()

	mockWord12.EXPECT().String().Return("12wxを").AnyTimes()
	mockWord12.EXPECT().StartTime().Return(d5s).AnyTimes()
	mockWord12.EXPECT().EndTime().Return(d6s).AnyTimes()

	// mockText.ReadWordのみで使う
	mockWordEOT.EXPECT().String().Return("").AnyTimes()

	mockText := mock_entity.NewMockText(c)
	gomock.InOrder(
		mockText.EXPECT().ReadWord().Return(mockWord1, true), //  1
		mockText.EXPECT().ReadWord().Return(mockWord2, true),
		mockText.EXPECT().ReadWord().Return(mockWord3, true),
		mockText.EXPECT().ReadWord().Return(mockWord4, true),
		mockText.EXPECT().ReadWord().Return(mockWord5, true), //  5
		mockText.EXPECT().ReadWord().Return(mockWord6, true),
		mockText.EXPECT().ReadWord().Return(mockWord7, true),
		mockText.EXPECT().ReadWord().Return(mockWord8, true),
		mockText.EXPECT().ReadWord().Return(mockWord9, true),
		mockText.EXPECT().ReadWord().Return(mockWord10, true), // 10
		mockText.EXPECT().ReadWord().Return(mockWord11, true),
		mockText.EXPECT().ReadWord().Return(mockWord12, true),   // 12
		mockText.EXPECT().ReadWord().Return(mockWordEOT, false), // EOT
	)

	// EOT専用
	mockTextEOT := mock_entity.NewMockText(c)

	mockConverter := mock_api.NewMockConverter(c)
	mockConverter.EXPECT().Send(gomock.Any).Return(nil).AnyTimes()
	mockConverter.EXPECT().CloseSend().Return(nil)
	gomock.InOrder(
		mockConverter.EXPECT().Recv(gomock.Any()).Return(mockText, nil),
		mockConverter.EXPECT().Recv(gomock.Any()).Return(mockTextEOT, io.EOF),
	)

	// 結果取得用chanel
	ch := make(chan string, 20)
	defer close(ch)
	mockIO := mock_presenter.NewMockIO(ch)

	// dummyのaudio file
	tmp, err := ioutil.TempFile("", "tmp.test.")
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}

	// create usecase
	keywords := []string{"1abい", "6klへ", "12wxを"}
	converter := mockConverter
	search := service.NewSearchWithWordsBeforeAndAfter(keywords, 5)
	pr := presenter.NewResultPresenter(mockIO)

	usecase := interactor.NewAudio2Text(converter, search, pr)

	// create controller
	cont := adapter.NewController(usecase)

	// run cmd
	go func() {
		cont.ConvertAudio2textFromFile(tmp.Name())
	}()

	result := make([]string, 0, 20)
	expected0 := "[pos: 1, start: 1s, end: 2s, keyword: 1abい] 1abい 2cdろ 3efは 4ghに 5ijほ 6klへ"
	expected1 := "[pos: 6, start: 3s, end: 4s, keyword: 6klへ] 1abい 2cdろ 3efは 4ghに 5ijほ 6klへ 7mnと 8opち 9qrり 10stぬ 11uvる"
	expected2 := "[pos: 12, start: 5s, end: 6s, keyword: 12wxを] 7mnと 8opち 9qrり 10stぬ 11uvる 12wxを"

	for {
		select {
		case s := <-ch:
			result = append(result, s)
		case <-time.After(50 * time.Millisecond):
			if result[0] != expected0 {
				t.Fatalf("main() return translated string: result[0]=%s, expected=%s", result[0], expected0)
			}
			if result[1] != expected1 {
				t.Fatalf("main() return translated string: result[1]=%s, expected=%s", result[1], expected1)
			}
			if result[2] != expected2 {
				t.Fatalf("main() return translated string: result[2]=%s, expected=%s", result[2], expected2)
			}
			return
		}
	}

}
