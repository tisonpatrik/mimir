package main

import (
	"mimir-scrapper/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/scrape", handlers.ScrapeHandler)
	http.ListenAndServe(":8080", nil)
}
