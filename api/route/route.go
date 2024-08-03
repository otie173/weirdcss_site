package route

import (
	"framework/api/handler"
	"net/http"
)

func Register() {
	http.HandleFunc("/", handler.HomepageHandler)
}
