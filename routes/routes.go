package routes

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SearchPlaylist(w http.ResponseWriter, req *http.Request) {
	// Retrieve request body
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalf("Unable to get request body %v", err.Error())
	}

	log.Printf("string %s", string(b))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `"{"alive": true}"`)
}
