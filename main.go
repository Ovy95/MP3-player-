package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Location and to the file you want to open has to be Mp3
	f, err := os.Open("./test.mp3")
	if err != nil {
		log.Fatal(err)

		// log.Print - Print to the stdout
		// os.Exit(1)

	}

	// we need to set up the streamer, format and check for the error
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalf("Error while decoding into mp3: %v\n", err)
	}
	defer streamer.Close()

	// speed/slow down the audio future features
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	//0 Tell speaker to play something
	//1 get something to play
	//2 stream it somewhere (streamer)
	//3 once it's done, then call the call back
	//4 say to the channel that the goroutine has finished
	<-done

}
