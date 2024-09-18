package api

import (
	"net/http"
)

func Register() {
	http.HandleFunc("/", contentHandler)
	http.HandleFunc("/weirdcss", contentHandler)
}
