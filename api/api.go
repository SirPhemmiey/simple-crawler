package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-crawler/crawler"
)

type SingleCrawlerRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type MultipleCrawlerRequest struct {
	URL []string `json:"urls" validate:"required,dive"`
}

func SingleCrawlerHandler(w http.ResponseWriter, r *http.Request) {

	request := r.Context().Value("validatedRequest").(*SingleCrawlerRequest)

	fmt.Println("SingleCrawlerHandler", r.Method, r.URL.Path, r.Body)

	content, err := crawler.SingleFetch(request.URL)
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

func MultipleCrawlerHandler(w http.ResponseWriter, r *http.Request) {

	request := r.Context().Value("validatedRequest").(*MultipleCrawlerRequest)

	fmt.Println("MultipleCrawlerHandler", r.Method, r.URL.Path, r.Body)

	content, err := crawler.MultipleFetch(request.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching URL: %v", err), http.StatusInternalServerError)
	}

	links, err := crawler.MultipleExtractLinks(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting links: %v", err), http.StatusInternalServerError)
	}
	var mappedLinks = make(map[string][]string)

	for index, _ := range request.URL {
		mappedLinks[request.URL[index]] = links[index]
	}

	response := struct {
		URL   []string            `json:urls`
		Links map[string][]string `json:links`
	}{
		URL:   request.URL,
		Links: mappedLinks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
