package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/huangc28/yt_music/routes"
)

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func newRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health", health)
	router.HandleFunc("/search-playlist", routes.SearchPlaylist)

	return router
}

func main() {
	routes := newRoutes()

	s := &http.Server{
		Addr:              ":8081",
		Handler:           routes,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
