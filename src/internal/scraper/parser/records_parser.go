package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

// ParseRecords parses and returns records and their associated events from the given document.
func ParseRecords(htmlContent string, sessionID uuid.UUID) ([]Record, error) {
	records := []Record{}

	parsedDocument, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	parsedDocument.Find(".stenovystoupeni").Each(func(index int, item *goquery.Selection) {
		// Extract the speaker name
		speakerElement := item.Find(".stenovystupujici a")
		speakerName := strings.TrimSpace(speakerElement.Text())

		// Remove the speaker reference element from the content
		speakerElement.Parent().Remove()

		// Extract and clean content
		content, err := item.Html()
		if err != nil {
			return
		}
		content = strings.TrimSpace(content)

		if speakerName != "" && content != "" {
			record := Record{
				SessionID:      sessionID,
				SpeakerName:    speakerName,
				Content:        content,
				Events:         extractEvents(item, speakerName),
				SequenceNumber: index + 1,
			}
			records = append(records, record)
		}
	})

	return records, nil
}

// extractEvents parses events from a speaker's content block, excluding links to the speaker's profile.
func extractEvents(item *goquery.Selection, speakerName string) []Event {
	events := []Event{}

	item.Find("a").Each(func(index int, link *goquery.Selection) {
		linkHref, exists := link.Attr("href")
		description := strings.TrimSpace(link.Text())

		// Exclude speaker profile links
		if exists && linkHref != "" && description != "" && !strings.Contains(description, speakerName) {
			event := Event{
				Link:        linkHref,
				Description: description,
			}
			events = append(events, event)
		}
	})

	return events
}
