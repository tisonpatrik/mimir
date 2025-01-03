package fetcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func FetchPage(url string) ([]Document, error) {
	c := colly.NewCollector()

	var documents []Document

	// Parse table rows
	c.OnHTML("table.PE_zebra > tbody > tr.body", func(e *colly.HTMLElement) {
		period, _ := e.DOM.Find("td:nth-of-type(1) id_obdobi").Html()
		meetingText := e.ChildText("td:nth-of-type(2) > a")
		meeting := 0
		if meetingText != "" {
			fmt.Sscanf(meetingText, "%d", &meeting)
		}
		document := e.ChildText("td:nth-of-type(6) > div")
		dateText := e.ChildText("td:nth-of-type(7)")
		date, _ := time.Parse("02.01.2006", dateText)
		docType := e.ChildText("td:nth-of-type(8) tdocname")
		originalLink := e.ChildAttr("td:nth-of-type(11) > a", "href")
		originalLink = e.Request.AbsoluteURL(originalLink)

		// Append the structured data
		documents = append(documents, Document{
			Period:       toInt(period),
			Meeting:      meeting,
			Document:     document,
			Date:         date,
			DocumentType: docType,
			OriginalLink: originalLink,
		})
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("failed to visit page: %w", err)
	}

	return documents, nil
}

// Helper function to convert string to int with default value 0
func toInt(value string) int {
	if value == "" {
		return 0
	}
	var result int
	fmt.Sscanf(strings.TrimSpace(value), "%d", &result)
	return result
}
