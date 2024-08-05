package handler

import (
	"framework/resources"
	"framework/utils/locate"
	"framework/utils/logger"
	"log"
	"net/http"
	"time"
)

var (
	templateRu []byte
	templateEn []byte
)

func init() {
	var err error
	templateRu, err = resources.GetTemplateFile("template/homepage_ru.html")
	if err != nil {
		log.Fatalf("Failed to load Russian template: %v", err)
	}
	templateEn, err = resources.GetTemplateFile("template/homepage_en.html")
	if err != nil {
		log.Fatalf("Failed to load English template: %v", err)
	}
}

func HomepageHandler(res http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		logger.LogRequest(req.Method, req.URL.Path, http.StatusMethodNotAllowed, time.Duration(time.Since(startTime).Microseconds()), req.RemoteAddr)
		return
	}

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		logger.LogRequest(req.Method, req.URL.Path, http.StatusNotFound, time.Duration(time.Since(startTime).Microseconds()), req.RemoteAddr)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	language := locate.GetLanguage(req)
	var template []byte
	if language == "ru" {
		template = templateRu
	} else {
		template = templateEn
	}

	_, err := res.Write(template)
	status := http.StatusOK
	if err != nil {
		log.Printf("Error writing response: %v", err)
		status = http.StatusInternalServerError
	}

	logger.LogRequest(req.Method, req.URL.Path, status, time.Duration(time.Since(startTime).Microseconds()), req.RemoteAddr)
}
