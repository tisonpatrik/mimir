package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"mimir-scrapper/src/internal/scraper/fetcher"
	"mimir-scrapper/src/internal/scraper/parser"
	"mimir-scrapper/src/pkg/utils"
	"net/http"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request, conn *pgxpool.Pool) {
	const (
		url       = "https://www.senat.cz/xqw/xervlet/pssenat/finddoc?typdok=steno"
		outputDir = "data/raw_data" // Directory for storing documents
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

	// Process and save documents
	parsedDocuments, err := processAndSaveDocuments(documents, outputDir)
	if err != nil {
		log.Println("Error processing documents:", err)
		http.Error(w, "Failed to process documents", http.StatusInternalServerError)
		return
	}

	// Respond with the number of successfully processed documents
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]int{"processed_documents": len(parsedDocuments)})
}

func processAndSaveDocuments(documents []string, outputDir string) ([]interface{}, error) {
	var parsedDocuments []interface{}

	for index, content := range documents {
		// Parse the HTML document
		transcript, err := parser.ParseHTMLDocument(content)
		if err != nil {
			log.Printf("Error parsing document %d: %v", index, err)
			continue
		}

		// Save the parsed transcript as a JSON file
		filename := filepath.Join(outputDir, fmt.Sprintf("document_%d.json", index+1))
		if err := utils.SaveToFile(filename, transcript); err != nil {
			log.Printf("Error saving document %d: %v", index, err)
			continue
		}

		parsedDocuments = append(parsedDocuments, transcript)
	}

	return parsedDocuments, nil
}
