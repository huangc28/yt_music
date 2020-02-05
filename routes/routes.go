package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/huangc28/yt_music/ytapi"
)

const (
	// YT_DEV_KEY youtube API developer key. @todo this key should be extract
	// to env file.
	YT_DEV_KEY = "AIzaSyD_E0RBoXGDvk68e7Dx-qwaW-RuOFptMzg"
)

type FailedResponse struct {
	Message string `json:"message"`
}

func SearchPlaylist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Retrieve request body
	query, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		log.Fatalf("unable to parse raw query: %s", err.Error())
	}

	if _, ok := query["search"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fr, _ := json.Marshal(&FailedResponse{
			Message: "query param 'search' is missing",
		})

		io.WriteString(w, string(fr))

		return
	}

	ytsrv, _ := ytapi.NewYoutubeAPI(YT_DEV_KEY)

	resp, err := ytsrv.SearchPlayList(query["search"][0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fr, _ := json.Marshal(&FailedResponse{
			Message: "Failed to request youtube API with the given creiteria",
		})

		io.WriteString(w, string(fr))

		return
	}

	log.Printf("DEBUG youtube resp %v", resp)

	// use search criteria to search for playlist item
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `"{"alive": true}"`)
}
