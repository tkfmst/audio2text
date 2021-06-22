package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	// 引数処理
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <AUDIOFILE>\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "<AUDIOFILE> must be a path to a local audio file. Audio file must be a 16-bit signed little-endian encoded with a sample rate of 16000.\n")
	}

	flag.Parse()

	// ここを2にして検索文字列を受け取る
	if len(flag.Args()) != 2 {
		log.Fatal("Please pass path to your local audio file as a command line argument")
	}
	audioFile := flag.Arg(0)
	log.Println("DEBUG", audioFile)

	searchWord := flag.Arg(1)
	log.Println("DEBUG", searchWord)

	// ファイル読み込み
	f, err := os.Open(audioFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("DEBUG", 1)

	// GCP setup config
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:              speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz:       8000,
					LanguageCode:          "ja-JP", // https://cloud.google.com/speech-to-text/docs/languages
					EnableWordTimeOffsets: true,
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	// request
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := f.Read(buf)
			log.Println("DEBUG", n)
			if n > 0 {
				log.Println(buf[:n])
				if err := stream.Send(&speechpb.StreamingRecognizeRequest{
					StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
						AudioContent: buf[:n],
					},
				}); err != nil {
					log.Printf("Could not send audio: %v", err)
				}
			}
			if err == io.EOF {
				log.Println("DEBUG", "EOF")
				// Nothing else to pipe, close the stream.
				if err := stream.CloseSend(); err != nil {
					log.Fatalf("Could not close stream: %v", err)
				}
				return
			}
			if err != nil {
				log.Printf("Could not read from %s: %v", audioFile, err)
				continue
			}
		}
	}()

	// 最初から11にしておいてshiftしてもエラーにしない
	result := make([]string, 11)

	// response process
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("DEBUG", "stream", "EOF")
			// FIXME 残ったWordを処理する
			break
		}
		if err != nil {
			log.Fatalf("Cannot stream results: %v", err)
		}
		if err := resp.Error; err != nil {
			log.Fatalf("Could not recognize: %v", err)
		}
		for _, res := range resp.Results {
			for _, alt := range res.Alternatives {
				for _, winfo := range alt.Words {
					fmt.Println("DEBUG", winfo.Word)
					fmt.Println("DEBUG", strings.Split(winfo.Word, "|")[0])

					result = result[1:]
					result = append(result, strings.Split(winfo.Word, "|")[0])

					fmt.Println("DEBUG", result[5])
					fmt.Println("DEBUG", len(result))
					if len(result) == 11 && result[5] == searchWord {
						log.Println("[Result]", "search=${searchWord}", strings.Trim(strings.Join(result, " "), " "))
					}
				}
			}

		}
	}
}
