package main

import (
	"bytes"
	"fmt"
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
	fmt.Println("Downloading video id:", id)
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	if err != nil {
		return nil, err
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats.Itag(140)[0])
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return &Video{Title: video.Title, Buffer: buf.Bytes()}, nil
}

func main() {
	if _, err := os.Stat("public/index.html"); os.IsNotExist(err) {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("public")))

	mux.HandleFunc("GET /dl", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET ", r.URL.RawPath)
		_, id, _ := strings.Cut(r.FormValue("url"), "v=")

		vid, err := DlSong(id)
		if err != nil {
			fmt.Println("Error from downloader:", err)
			http.Error(w, "Something went wrong while downloading this song. Please try again", 500)
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

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, mux)
}
