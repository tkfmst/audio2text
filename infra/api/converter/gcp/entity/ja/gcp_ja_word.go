package ja

import (
	"strings"
	"time"

	"example.com/audio2text/entity"
	"example.com/audio2text/entity/id"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

// GCPJaWord is entity.Word implementation
type GCPJaWord struct {
	*id.WordID
	tid       *id.TextID
	raw       *speechpb.WordInfo
	word      string
	howToRead string
	startTime *durationpb.Duration
	endTime   *durationpb.Duration
}

func NewGCPJaWord(tid *id.TextID, raw *speechpb.WordInfo) entity.Word {
	w := new(GCPJaWord)
	w.WordID = id.NewWordID()
	w.raw = raw
	w.word, w.howToRead = splitWordAndHowToRead(raw)
	w.tid = tid
	w.startTime = raw.StartTime
	w.endTime = raw.EndTime

	return w
}

func (w *GCPJaWord) TextID() *id.TextID {
	return w.tid
}

func (w *GCPJaWord) String() string {
	return w.word
}

func (w *GCPJaWord) StartTime() time.Duration {
	return w.startTime.AsDuration()
}

func (w *GCPJaWord) EndTime() time.Duration {
	return w.endTime.AsDuration()
}

// jaでは`単語|読み`というフォーマットの文字列が返ってくる
func splitWordAndHowToRead(raw *speechpb.WordInfo) (string, string) {
	arr := strings.Split(raw.Word, "|")
	return strings.Trim(arr[0], " "), arr[1]
}
