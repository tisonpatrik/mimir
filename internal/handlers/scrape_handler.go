package handlers

import (
	"fmt"
	"net/http"

	"mimir-scrapper/internal/fetcher"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	const url = "https://www.senat.cz/xqw/xervlet/pssenat/finddoc?typdok=steno"
	html, err := fetcher.FetchPage(url)
	if err != nil {
		http.Error(w, "Failed to fetch page", http.StatusInternalServerError)
		return
	}
	fmt.Print(html)
}
