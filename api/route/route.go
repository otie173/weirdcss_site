package route

import (
	"framework/api/handler"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", handler.HomepageHandler)
}
