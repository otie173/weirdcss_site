package locate

import (
	"net/http"
	"strings"
)

func GetLanguage(req *http.Request) string {
	language := req.Header.Get("Accept-Language")
	if strings.Contains(language, "ru") {
		return "ru"
	}
	return "en"
}
