package fetcher

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func FetchPage(url string) ([]string, error) {
	c := colly.NewCollector()

	var documents []string

	// Callback for finding "Originál" links
	c.OnHTML("a[title='Otevře originální dokument ']", func(e *colly.HTMLElement) {
		// Build the full URL (assuming relative links)
		link := e.Request.AbsoluteURL(e.Attr("href"))

		// Visit and read the linked page
		c.Visit(link)
	})

	// Callback for visiting each "originál" link and saving its content
	c.OnResponse(func(r *colly.Response) {
		if r.Request.URL.String() != url { // Ensure we don't add the main URL content
			content := string(r.Body)
			documents = append(documents, content) // Save the content
		}
	})

	// Start scraping only the provided URL for links, no direct content fetch
	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("failed to visit main page: %w", err)
	}

	return documents, nil
}
