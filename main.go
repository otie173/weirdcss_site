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

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
