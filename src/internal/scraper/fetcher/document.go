package fetcher

import "time"

type Document struct {
	Period       int       `json:"period"`
	Meeting      int       `json:"meeting"`
	Document     string    `json:"document"`
	Date         time.Time `json:"date"`
	DocumentType string    `json:"document_type"`
	OriginalLink string    `json:"original_link"`
}
