package main

import (
	"log"
	"net/http"

	"framework/api/route"
	"framework/utils/logger"
)

func main() {
	logger.Register()
	route.Register()

	// Created file server for css file, which used in html templates
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("static/", http.StripPrefix("static", fileServer))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
