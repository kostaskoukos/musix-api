package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from golang"))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
