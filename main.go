package main

import (
	"net/http"
)

func main() {
	// Created file server for css file, which used in html templates
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("static/", http.StripPrefix("static", fileServer))
}
