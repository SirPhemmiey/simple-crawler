package main

import (
	"log"
	"net/http"
	"web-crawler/api"
)

func main() {
	http.HandleFunc("POST /crawl", api.CrawlerHandler)

	log.Println("Starting server on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
