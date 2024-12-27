package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

type Transcript struct {
	Title   string
	Entries []Entry
}

type Entry struct {
	Speaker     string
	SpeakerLink string
	Text        string
	Events      []Event
}

type Event struct {
	Description string
	Link        string
}

func ParseHTMLDocument(html string) (*Transcript, error) {
	// Decode HTML from Windows-1250 to UTF-8
	decoder := charmap.Windows1250.NewDecoder()
	decodedHTML, err := decoder.String(html)
	if err != nil {
		return nil, err
	}

	// Load document into goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(decodedHTML))
	if err != nil {
		return nil, err
	}

	// Extract title
	title := doc.Find("title").Text()

	// Initialize transcript
	transcript := &Transcript{Title: title}

	// Parse speech sections
	doc.Find("div.stenovystoupeni").Each(func(i int, s *goquery.Selection) {
		// Find speaker
		speakerLink := s.Find("span.stenovystupujici a")
		speaker := speakerLink.Text()
		speaker = strings.TrimSpace(speaker)

		// Extract speaker link
		speakerHref, _ := speakerLink.Attr("href")

		// Find speech text
		text := s.Find("p").Text()
		text = strings.TrimSpace(text)

		// Extract events
		var events []Event
		s.Find("a[href]").Each(func(i int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if exists && !strings.Contains(href, "webtv/archiv") { // Ignore audio/video links
				description := link.Text()
				events = append(events, Event{
					Description: strings.TrimSpace(description),
					Link:        href,
				})
			}
		})

		// Add entry to transcript
		if speaker != "" && text != "" {
			transcript.Entries = append(transcript.Entries, Entry{
				Speaker:     speaker,
				SpeakerLink: speakerHref,
				Text:        text,
				Events:      events,
			})
		}
	})

	return transcript, nil
}
