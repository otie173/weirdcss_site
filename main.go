package main

import (
	"io/fs"
	"log"
	"net/http"

	"framework/api/route"
	"framework/resources"
	"framework/utils/logger"
)

func main() {
	logger.Register()
	route.Register()

	staticFS, err := fs.Sub(resources.StaticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Created file server for css file, which used in html templates
	fileServer := http.FileServer(http.FS(staticFS))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
