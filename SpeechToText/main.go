// export GOOGLE_APPLICATION_CREDENTIALS=credentials.json

// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command caption sends audio data to the Google Speech API
// and prints its transcript.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

const usage = `Usage: caption <audiofile>

Audio file must be a 16-bit signed little-endian encoded
with a sample rate of 16000.

The path to the audio file may be a GCS URI (gs://...).
`

func main() {
  fmt.Println("m0")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

  fmt.Println("m1")
	var runFunc func(io.Writer, string) error
  fmt.Println("m2")
	path := os.Args[1]
	if strings.Contains(path, "://") {
		runFunc = recognizeGCS
	} else {
		runFunc = recognize
	}
  fmt.Println("m3")

	// Perform the request.
	if err := runFunc(os.Stdout, os.Args[1]); err != nil {
    fmt.Println("err" + err.Error())
		log.Fatal(err)
	}
  fmt.Println("m4")
}

// [START speech_transcribe_sync_gcs]

func recognizeGCS(w io.Writer, gcsURI string) error {
  fmt.Println("1")
	ctx := context.Background()
fmt.Println("2")
	client, err := speech.NewClient(ctx)
	if err != nil {
		return err
	}
fmt.Println("3")
	// Send the request with the URI (gs://...)
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			// Encoding:        speechpb.RecognitionConfig_FLAC,
			// SampleRateHertz: 16000,
			LanguageCode:    "pl-PL",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: gcsURI},
		},
	})

  fmt.Println("results")

	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return nil
}

// [END speech_transcribe_sync_gcs]

// [START speech_transcribe_sync]

func recognize(w io.Writer, file string) error {
  fmt.Println("r1")
	ctx := context.Background()
  fmt.Println("r2")
	client, err := speech.NewClient(ctx)
	if err != nil {
		return err
	}
  fmt.Println("r3")

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
  fmt.Println("r4")

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			LanguageCode:    "pl-PL",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
  fmt.Println("r5")
  if err != nil {
		return err
	}
  fmt.Println("r6")
  fmt.Println(resp.Results)

	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return nil
}

// [END speech_transcribe_sync]
