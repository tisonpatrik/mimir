package scraper

import (
	"context"
	"encoding/json"
	"log"
	"mimir-scrapper/src/internal/scraper/fetcher"
	"mimir-scrapper/src/internal/scraper/parser"
	"mimir-scrapper/src/pkg/repository"
	"mimir-scrapper/src/pkg/services"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, pool *pgxpool.Pool) {
	const (
		url       = "https://www.senat.cz/xqw/xervlet/pssenat/finddoc?typdok=steno"
		outputDir = "data/raw_data" // Directory for storing documents
	)

	// Fetch structured documents
	documents, err := fetcher.FetchPage(url)
	if err != nil {
		log.Println("Error fetching documents:", err)
		http.Error(w, "Failed to fetch documents", http.StatusInternalServerError)
		return
	}

	// Process and save documents
	parsedDocuments, err := processAndSaveDocuments(ctx, pool, documents)
	if err != nil {
		log.Println("Error processing documents:", err)
		http.Error(w, "Failed to process documents", http.StatusInternalServerError)
		return
	}

	// Respond with the number of successfully processed documents
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]int{"processed_documents": len(parsedDocuments)})
}

func processAndSaveDocuments(ctx context.Context, pool *pgxpool.Pool, documents []fetcher.Document) ([]interface{}, error) {
	institutionName := "Senát"
	occasionName := "meeting"
	var parsedDocuments []interface{}

	// Start transaction
	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Initialize sqlc Queries and SessionService
	queries := repository.New(tx)
	sessionService := services.NewSessionService(queries)

	// Process documents
	for _, doc := range documents {
		// Get or create session
		session, err := sessionService.GetOrCreateSession(ctx, institutionName, occasionName, doc.Date)
		if err != nil {
			log.Printf("Error creating session: %v", err)
			return nil, err
		}

		// Parse HTML document for speakers
		speakers, err := parser.ParseSpeakers(doc.HTMLContent)
		if err != nil {
			log.Printf("Error parsing speakers for document dated %s: %v", doc.Date, err)
			return nil, err
		}

		// Parse HTML document for records and events
		records, err := parser.ParseRecords(doc.HTMLContent, session.ID)
		if err != nil {
			log.Printf("Error parsing records for document dated %s: %v", doc.Date, err)
			return nil, err
		}
		if len(speakers) > 0 {
			log.Printf("First Speaker: %+v", speakers[0])
		} else {
			log.Println("No speakers found.")
		}

		if len(records) > 0 {
			log.Printf("First Record: %+v", records[0])
		} else {
			log.Println("No records found.")
		}

	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return nil, err
	}

	return parsedDocuments, nil
}
