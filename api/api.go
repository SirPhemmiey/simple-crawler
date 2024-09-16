package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-crawler/crawler"
)

func CrawlerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var request struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

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
