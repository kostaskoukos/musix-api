package main

import (
	"io"
	"net/http"
	"os"

	"github.com/kkdai/youtube/v2"
)

func DlSong(id string) {
	videoID := "BaW_jenozKc"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create("video.mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		DlSong("")
		w.Write([]byte("hi from golang"))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
