package handler

import (
	"fmt"
	"net/http"
)

func HomepageHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello homepage!")
}
