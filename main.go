package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
)

type Video struct {
	Title  string
	Buffer []byte
}

var (
	jar, _     = cookiejar.New(nil)
	httpClient = &http.Client{Jar: jar}
	client     = &youtube.Client{HTTPClient: httpClient}
	cookies    = []*http.Cookie{
		{
			Name:     "__Secure-1PAPISID",
			Value:    "z8VfquYWmxJvYKb8/Ag2BvzH6lOQCa8D-_",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   true,
			HttpOnly: false,
		},
		{
			Name:     "__Secure-1PSID",
			Value:    "g.a000nQhaeac0OY1WkLwf-k9vjw2YP5xEObHwu6b8RXxwTTfEzjzZ11jzGFPmnjDQfhcYmjsWNAACgYKAcESARASFQHGX2MiQWySc63JUkHffUAQvdy7fhoVAUF8yKoB9OWzR_SWztFwV-o6B2IW0076",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-1PSIDCC",
			Value:    "AKEyXzU1AKXlYmB_GReYmql-8-8jS3S3BADov5x4jq6pbxQOXXqNDTz8CxxTr3FQJjrsCemN4w",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1756126433, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-1PSIDTS",
			Value:    "sidts-CjIBUFGohxxQh2QPrR65gt6dLbA01wwEsKV3soJiBQovgoBPx0mFfF2GL-x_aMHIgxkYGxAA",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1756126422, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-3PAPISID",
			Value:    "z8VfquYWmxJvYKb8/Ag2BvzH6lOQCa8D-_",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   true,
			HttpOnly: false,
		},
		{
			Name:     "__Secure-3PSID",
			Value:    "g.a000nQhaeac0OY1WkLwf-k9vjw2YP5xEObHwu6b8RXxwTTfEzjzZPTuGkkJxkIFHQDgEaVPSEgACgYKAekSARASFQHGX2MinTKVac9pkIDNqOml-SlmgBoVAUF8yKomb6Gx655_VP65zvlFeiwZ0076",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-3PSIDCC",
			Value:    "AKEyXzXnW88mzJpYxB4bseKfCSflX0Dxy7lvvv6oZevXQBZ-Qxwp4J7sccCRN5vLQwJ9FsWl",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1756126433, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-3PSIDTS",
			Value:    "sidts-CjIBUFGohxxQh2QPrR65gt6dLbA01wwEsKV3soJiBQovgoBPx0mFfF2GL-x_aMHIgxkYGxAA",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1756126422, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "__Secure-YEC",
			Value:    "CgthYzd6NXdPZU8tayjg0qy2BjIiCgJHUhIcEhgSFhMLFBUWFwwYGRobHB0eHw4PIBAREiEgGA%3D%3D",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1758718163, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "APISID",
			Value:    "HxXXy8TYEAAt-ict/AKKTqlP5sEFqjx6XT",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   false,
			HttpOnly: false,
		},
		{
			Name:     "CONSISTENCY",
			Value:    "AKreu9tlZKG1nizDSuavser5hmQekB2ropGx_606heOJpotOWz0GSFM4Sa1Fj8dQAf9sVDA5upmGdQPo4Q68VxoTdiYCKdXB2s1z3JnxIe8mm2hj_DHzAVmWtW18N71qli-IfvV-0SCUTYHTUekrZPKF",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1724590954, 0),
			Secure:   true,
			HttpOnly: false,
		},
		{
			Name:     "HSID",
			Value:    "AGEUqqoo2JuhMyVkq",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   false,
			HttpOnly: true,
		},
		{
			Name:     "LOGIN_INFO",
			Value:    "AFmmF2swRQIgcPqBxcWQ1Ynanr10srMpMQLjf_s3PqEZm_Wm1poiiJICIQDbzSu2B7gMzMNtO64N4-4GxNUAhguYj-T5e7WzzhWv7Q:QUQ3MjNmeEJLWWI5bG8yZlhvTEhtSl9remtzei1mQ2V3WEtnSnQzS2dRaUtlTVloRXdxN1NjU2tzSERpc3M1eFN1Z1pKTFpZRTJqN0FRWldhRWZGSExwS3RiZ0ZXZjVNNUp4djlKTmZsLVFYZG1yZERJaFVIZmFJM0VwNXFXbl9fb0Iya0VqdW1CUW10dXRpajhQZEk1ZjAyd3M1eTAwUTRn",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150423, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "NID",
			Value:    "517=ssdHu6JV7lPH8r2IGLH6xQVymbCLRDsrgwPoaFt-AgqhVAFbK1GgkHSGpgmreGN-Z3L7t4xg2nE7ABi1cgE4wg-pUhM-Bd40d_TVx1buO0bAjXamMYBTCB-GX93rFVIyToojSpG8Vvk8c10gWd1qR0wxhrP3ZHHs37d9PqVH9kEa2Quzw9udRv8K5E4pnXKDcA",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1740401552, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "SAPISID",
			Value:    "FWz2k9Oh1BQWkKfH/AcQ1pJrS91KFDpxAA",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   false,
			HttpOnly: false,
		},
		{
			Name:     "SEARCH_SAMESITE",
			Value:    "CgQIoAI=",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1727218890, 0),
			Secure:   false,
			HttpOnly: false,
		},
		{
			Name:     "SID",
			Value:    "Twj1w3wJDyiZNLKr4K4XaIHrrtXUHdJdnLPxxDE2zJrLM54Xk5QnIIXke4AQXl7uH0MP6w.",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   false,
			HttpOnly: true,
		},
		{
			Name:     "SIDCC",
			Value:    "AJi4QfGnxBGCwZD3DCR9cbsaYOENzMgAFvCjKzEFhT1naBPprAvpM7a6DlRks2n1yG1ZzJBz9A",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1756126433, 0),
			Secure:   true,
			HttpOnly: true,
		},
		{
			Name:     "SSID",
			Value:    "AF2UvoKD4FyFIC3B2",
			Path:     "/",
			Domain:   "https://youtube.com",
			Expires:  time.Unix(1759150422, 0),
			Secure:   false,
			HttpOnly: true,
		},
	}
)

func DlSong(id string) (*Video, error) {
	fmt.Println("Downloading video id:", id)

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
	url, _ := url.Parse("https://youtube.com")
	httpClient.Jar.SetCookies(url, cookies)

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
