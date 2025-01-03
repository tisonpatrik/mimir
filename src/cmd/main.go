package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"mimir-scrapper/src/internal/scraper"
	"mimir-scrapper/src/pkg/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	// Inicializace connection poolu pomocí db balíčku
	pool, err := db.NewPool(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		scraper.ScrapeHandler(w, r, pool)
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
