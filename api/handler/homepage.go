package handler

import (
	"framework/utils/logger"
	"log"
	"net/http"
	"os"
	"time"
)

func HomepageHandler(res http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	htmlTemplate, err := os.ReadFile("templates/homepage.html")
	if err != nil {
		log.Fatalln(err)
	}

	res.Header().Set("Content-Type", "text/html")
	res.Write(htmlTemplate)

	logger.LogRequest(req.Method, req.URL.Path, http.StatusOK, time.Duration(time.Since(startTime).Microseconds()))
}
