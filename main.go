package main

import (
	"log"
	"net/http"
	"web-crawler/api"
	"web-crawler/api/middleware"
)

func main() {

	SetupRoute()
	log.Println("Starting server on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func SetupRoute() {
	http.HandleFunc("POST /crawl", middleware.ValidateRequest(api.CrawlerRequest{})(api.CrawlerHandler))

}
