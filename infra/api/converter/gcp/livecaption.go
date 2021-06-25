package gcp

import (
	"context"
	"fmt"
	"io"

	speech "cloud.google.com/go/speech/apiv1"
	"example.com/audio2text/app/if/api"
	"example.com/audio2text/entity"
	"example.com/audio2text/entity/enum"
	"example.com/audio2text/entity/id"
	"example.com/audio2text/infra/api/converter/gcp/entity/ja"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type Livecaption struct {
	stream speechpb.Speech_StreamingRecognizeClient
}

func NewLivecaption(lang enum.Lang, sRate int32) (api.Converter, error) {
	// GCP setup config
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:              speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz:       sRate,
					LanguageCode:          lang.String(),
					EnableWordTimeOffsets: true,
				},
			},
		},
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Livecaption{stream}, nil
}

func (lcap *Livecaption) Send(buf []byte) error {
	if err := lcap.stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
			AudioContent: buf,
		},
	}); err != nil {
		return errors.Wrap(err, "could not send audio")
	}

	return nil
}

func (lcap *Livecaption) CloseSend() error {
	if err := lcap.stream.CloseSend(); err != nil {
		return errors.Wrap(err, "could not close stream")
	}
	return nil
}

func (lcap *Livecaption) Recv(aid *id.AudioID) (entity.Text, error) {
	resp, err := lcap.stream.Recv()
	if err == io.EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot stream result: %v", err)
	}
	if err := resp.Error; err != nil {
		return nil, errors.WithStack(fmt.Errorf("could not recognize: %v", err))
	}

	return ja.NewGCPJaText(aid, resp), nil
}
