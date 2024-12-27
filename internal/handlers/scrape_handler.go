package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"mimir-scrapper/internal/fetcher"
	"mimir-scrapper/internal/parser"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	const url = "https://www.senat.cz/xqw/xervlet/pssenat/finddoc?typdok=steno"

	// Fetch HTML documents
	documents, err := fetcher.FetchPage(url)
	if err != nil {
		log.Println("Error fetching documents:", err)
		http.Error(w, "Failed to fetch documents", http.StatusInternalServerError)
		return
	}

	// Parse the first document only
	if len(documents) > 0 {
		transcript, err := parser.ParseHTMLDocument(documents[0])
		if err != nil {
			log.Println("Error parsing document:", err)
			http.Error(w, "Failed to parse document", http.StatusInternalServerError)
			return
		}

		// Return the parsed transcript as JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(transcript)
	} else {
		http.Error(w, "No documents found", http.StatusNotFound)
	}
}
