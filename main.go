package main

import (
	"log"
	"net/http"
	"site/api"
	"site/utils/logger"
)

func main() {
	logger.Register()
	api.Register()

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
