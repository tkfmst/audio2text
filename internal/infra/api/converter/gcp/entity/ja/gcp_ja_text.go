package ja

import (
	"example.com/audio2text/internal/entity"
	"example.com/audio2text/internal/entity/id"
	"example.com/audio2text/internal/util/logger"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// GCPJaText is entity.Text implementation
type GCPJaText struct {
	*id.TextID
	audioID *id.AudioID
	ch      chan entity.Word
}

func NewGCPJaText(aid *id.AudioID, raw *speechpb.StreamingRecognizeResponse) entity.Text {
	tid := id.NewTextID()

	ch := make(chan entity.Word)
	go func() {
		for _, res := range raw.Results {
			logger.Debug("gcp response result", res)
			for _, alt := range res.Alternatives {
				logger.Debug("word count", len(alt.Words))
				for _, wInfo := range alt.Words {
					ch <- NewGCPJaWord(tid, wInfo)
				}
			}
		}
		close(ch)
	}()

	t := new(GCPJaText)
	t.TextID = tid
	t.audioID = aid
	t.ch = ch

	return t
}

func (t *GCPJaText) ReadWord() (entity.Word, bool) {
	w, ok := <-t.ch
	return w, ok
}

func (t *GCPJaText) AudioID() *id.AudioID {
	return t.audioID
}
