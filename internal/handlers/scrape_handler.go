package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"mimir-scrapper/internal/fetcher"
	"mimir-scrapper/internal/parser"
	"mimir-scrapper/pkg/utils"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	const (
		url       = "https://www.senat.cz/xqw/xervlet/pssenat/finddoc?typdok=steno"
		outputDir = "documents" // Directory for storing documents
	)

	// Ensure the output directory exists
	if err := utils.EnsureDir(outputDir); err != nil {
		log.Println("Error creating output directory:", err)
		http.Error(w, "Failed to set up storage", http.StatusInternalServerError)
		return
	}

	// Fetch HTML documents
	documents, err := fetcher.FetchPage(url)
	if err != nil {
		log.Println("Error fetching documents:", err)
		http.Error(w, "Failed to fetch documents", http.StatusInternalServerError)
		return
	}

	// Process each document
	var parsedDocuments []interface{}
	for i, doc := range documents {
		transcript, err := parser.ParseHTMLDocument(doc)
		if err != nil {
			log.Printf("Error parsing document %d: %v", i, err)
			continue // Skip this document and proceed to the next
		}

		// Save the parsed transcript as a JSON file
		filename := filepath.Join(outputDir, fmt.Sprintf("document_%d.json", i+1))
		if err := utils.SaveToFile(filename, transcript); err != nil {
			log.Printf("Error saving document %d: %v", i, err)
			continue
		}

		parsedDocuments = append(parsedDocuments, transcript)
	}

	// Respond with the number of successfully processed documents
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]int{"processed_documents": len(parsedDocuments)})
}
