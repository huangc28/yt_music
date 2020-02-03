package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/huangc28/yt_music/ytapi"
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

	// @todo, if search query param doesn't contain "search" echo error.
	ytDevKey := "AIzaSyD_E0RBoXGDvk68e7Dx-qwaW-RuOFptMzg"
	ytapi.NewYoutubeAPI(ytDevKey)

	// use search criteria to search for playlist item

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `"{"alive": true}"`)
}
