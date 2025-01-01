package main

import (
	"net/http"

	"mimir-scrapper/src/internal/handlers"
)

func main() {
	http.HandleFunc("/scrape", handlers.ScrapeHandler)
	http.ListenAndServe(":8080", nil)
}
