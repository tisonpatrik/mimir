package handlers

import "net/http"

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Scrape handler is working!"))
}
