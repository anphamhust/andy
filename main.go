package main

// #cgo LDFLAGS: -lasound
// #include "arecord.h"
import "C"

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	fmt.Println("Go Vision!")
	fileName := "tmp.wav"
	cfileName := C.CString(fileName)
	C.record_to_file(cfileName, C.int(3))
	recognize(os.Stdout, fileName)
}

func recognize(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return nil
}
