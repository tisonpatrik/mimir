package fetcher

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func FetchPage(url string) (string, error) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return string(url), nil
}
