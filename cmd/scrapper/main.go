package main

import (
	"net/http"

	"mimir-scrapper/internal/handlers"
)

func main() {
	http.HandleFunc("/scrape", handlers.ScrapeHandler)
	http.ListenAndServe(":8080", nil)
}
