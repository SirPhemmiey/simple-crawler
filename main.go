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
	http.HandleFunc("POST /crawl-single", middleware.ValidateRequest(api.SingleCrawlerRequest{})(api.SingleCrawlerHandler))
	http.HandleFunc("POST /crawl-multiple", middleware.ValidateRequest(api.MultipleCrawlerRequest{})(api.MultipleCrawlerHandler))

}
