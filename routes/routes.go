package routes

import (
	"fmt"
	"net/http"
)

func SearchPlaylist(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "playlists\n")
}
