package main

import (
	"bytes"
	"net/http"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

type Video struct {
	Title  string
	Buffer []byte
}

func DlSong(id string) (*Video, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	if err != nil {
		return nil, err
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return &Video{Title: video.Title, Buffer: buf.Bytes()}, nil
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("public")))

	mux.HandleFunc("GET /dl", func(w http.ResponseWriter, r *http.Request) {
		_, id, _ := strings.Cut(r.FormValue("url"), "v=")

		vid, err := DlSong(id)
		if err != nil {
			http.Error(w, "Something went wrong while donwloading this song. Please try again", 500)
			return
		}

		w.Header().Set("Content-Disposition", `attachment; filename="`+vid.Title+`.mp3"`)
		w.Header().Set("Content-Type", "audio/mp3")
		w.Write(vid.Buffer)
	})

	port := ":"
	if p, _ := os.LookupEnv("PORT"); p != "" {
		port += p
	} else {
		port += "3000"
	}

	http.ListenAndServe(port, mux)
}
