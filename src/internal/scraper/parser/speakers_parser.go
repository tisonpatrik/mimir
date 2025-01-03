package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseSpeakers extracts and returns a unique list of speakers from the given document.
func ParseSpeakers(htmlContent string) ([]Speaker, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	speakerSet := make(map[string]bool)
	var speakers []Speaker

	doc.Find(".stenovystupujici a").Each(func(index int, item *goquery.Selection) {
		name := strings.TrimSpace(item.Text())
		if name != "" && !speakerSet[name] {
			speakerSet[name] = true
			speakers = append(speakers, Speaker{Name: name})
		}
	})

	return speakers, nil
}
