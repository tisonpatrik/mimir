package fetcher

import (
	"io"
	"log"
	"mimir-scrapper/src/pkg/utils"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func FetchPage(url string) ([]Document, error) {
	c := colly.NewCollector()
	var documents []Document

	// Parse table rows
	c.OnHTML("table.PE_zebra > tbody > tr.body", func(e *colly.HTMLElement) {
		parseTableRow(e, &documents)
	})

	// Start scraping
	if err := c.Visit(url); err != nil {
		log.Printf("Error visiting URL: %s, Error: %v", url, err)
		return nil, err
	}

	return documents, nil
}

func parseTableRow(e *colly.HTMLElement, documents *[]Document) {
	period := utils.ToInt(strings.TrimSpace(e.ChildText("td:nth-of-type(1)")))
	meeting := utils.ToInt(strings.TrimSpace(e.ChildText("td:nth-of-type(2) > a")))
	document := strings.TrimSpace(e.ChildText("td:nth-of-type(6) div"))
	date, _ := time.Parse("02.01.2006", strings.TrimSpace(e.ChildText("td:nth-of-type(7)")))
	docType := strings.TrimSpace(e.ChildText("td:nth-of-type(8)"))
	originalLink := e.ChildAttr("td:nth-of-type(11) a.icon.iconBefore.html", "href")

	if originalLink != "" {
		originalLink = e.Request.AbsoluteURL(originalLink)
	}

	doc := Document{
		Period:       period,
		Meeting:      meeting,
		Document:     document,
		Date:         date,
		DocumentType: docType,
		OriginalLink: originalLink,
	}

	// Fetch HTML content
	if originalLink != "" {
		htmlContent := fetchHTMLContent(originalLink)

		doc.HTMLContent = htmlContent
	}

	*documents = append(*documents, doc)
}

// fetchHTMLContent fetches HTML content for a given URL
func fetchHTMLContent(url string) string {
	htmlContent := ""
	collector := colly.NewCollector()

	collector.OnResponse(func(r *colly.Response) {
		// Decode windows-1250 content to UTF-8
		reader := transform.NewReader(strings.NewReader(string(r.Body)), charmap.Windows1250.NewDecoder())
		decoded, err := io.ReadAll(reader)
		if err != nil {
			log.Printf("Error decoding content from URL: %s, Error: %v", url, err)
			htmlContent = "Error decoding content: " + err.Error()
			return
		}
		htmlContent = string(decoded)
	})

	if err := collector.Visit(url); err != nil {
		log.Printf("Error fetching content from URL: %s, Error: %v", url, err)
		return "Error fetching content: " + err.Error()
	}
	return htmlContent
}
