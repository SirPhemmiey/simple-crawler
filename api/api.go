package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-crawler/crawler"
)

type CrawlerRequest struct {
	URL string `json:"url" validate:"required,url"`
}

func CrawlerHandler(w http.ResponseWriter, r *http.Request) {

	request := r.Context().Value("validatedRequest").(*CrawlerRequest)

	fmt.Println("CrawlerHandler", r.Method, r.URL.Path, r.Body)

	content, err := crawler.Fetch(request.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching URL: %v", err), http.StatusInternalServerError)
	}

	links, err := crawler.ExtractLinks(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting links: %v", err), http.StatusInternalServerError)
	}

	response := struct {
		URL   string   `json:url`
		Links []string `json:links`
	}{
		URL:   request.URL,
		Links: links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
