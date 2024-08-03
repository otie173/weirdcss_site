package main

import (
	"log"
	"net/http"

	"framework/api/route"
	"framework/utils/logger"
)

func main() {
	mux := http.NewServeMux()

	logger.Register()
	route.Register(mux)

	// Created file server for css file, which used in html templates
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalln(err)
	}
}
