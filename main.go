package main

import (
	"bytes"
	"net/http"
	"os"

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

	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		vid, err := DlSong(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Disposition", `attachment; filename="`+vid.Title+`".mp3`)
		w.Header().Set("Content-Type", "audio/mp3")
		w.Write(vid.Buffer)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
