package handler

import (
	"fmt"
	"framework/utils/logger"
	"net/http"
	"time"
)

func HomepageHandler(res http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	fmt.Fprint(res, "Hello homepage!")

	logger.LogRequest(req.Method, req.URL.Path, http.StatusOK, time.Duration(time.Since(startTime).Microseconds()))
}
