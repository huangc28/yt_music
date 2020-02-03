package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchPlayListRoute(t *testing.T) {
	// initialize a new request to search playlist API
	req, err := http.NewRequest("GET", "/search-playlist", nil)
	q := req.URL.Query()
	q.Add("search", "hiphop")
	req.URL.RawQuery = q.Encode()

	log.Printf("query string %v", req.URL.String())

	if err != nil {
		log.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchPlaylist)

	handler.ServeHTTP(rr, req)

	log.Printf("status code %v", rr.Code)
	log.Printf("Body %v", rr.Body.String())
}
