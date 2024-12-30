package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"

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

	results := make(chan Result, len(documents))
	var wg sync.WaitGroup

	for i, doc := range documents {
		wg.Add(1)
		go func(index int, content string) {
			defer wg.Done()
			transcript, err := parser.ParseHTMLDocument(content)
			results <- Result{Index: index, Transcript: transcript, Error: err}
		}(i, doc)
	}

	// Close results channel once all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and save parsed transcripts
	var parsedDocuments []interface{}
	for result := range results {
		if result.Error != nil {
			log.Printf("Error parsing document %d: %v", result.Index, result.Error)
			continue
		}

		// Save the parsed transcript as a JSON file
		filename := filepath.Join(outputDir, fmt.Sprintf("document_%d.json", result.Index+1))
		if err := utils.SaveToFile(filename, result.Transcript); err != nil {
			log.Printf("Error saving document %d: %v", result.Index, err)
			continue
		}

		parsedDocuments = append(parsedDocuments, result.Transcript)
	}

	// Respond with the number of successfully processed documents
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]int{"processed_documents": len(parsedDocuments)})
}
