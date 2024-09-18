package api

import (
	"log"
	"net/http"
	"site/resource"
	"site/utils/locate"
)

var (
	aboutTemplateRu     []byte
	aboutTemplateEn     []byte
	frameworkTemplateRu []byte
	frameworkTemplateEn []byte
)

func init() {
	var err error
	frameworkTemplateRu, err = resource.GetTemplateFile("template/weirdcss/weirdcss_ru.html")
	if err != nil {
		log.Fatalf("Failed to load Russian template: %v", err)
	}
	frameworkTemplateEn, err = resource.GetTemplateFile("template/weirdcss/weirdcss_en.html")
	if err != nil {
		log.Fatalf("Failed to load English template: %v", err)
	}
	aboutTemplateRu, err = resource.GetTemplateFile("template/otie173/otie173_ru.html")
	if err != nil {
		log.Fatalf("Failed to load English template: %v", err)
	}
	aboutTemplateEn, err = resource.GetTemplateFile("template/otie173/otie173_en.html")
	if err != nil {
		log.Fatalf("Failed to load English template: %v", err)
	}

}

func contentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Error: Method not allowed - %s %s", req.Method, req.URL.Path)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	language := locate.GetLanguage(req)

	var template []byte
	switch req.URL.Path {
	case "/":
		if language == "ru" {
			template = aboutTemplateRu
		} else {
			template = aboutTemplateEn
		}
	case "/weirdcss":
		if language == "ru" {
			template = frameworkTemplateRu
		} else {
			template = frameworkTemplateEn
		}
	}

	if _, err := res.Write(template); err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}
}
