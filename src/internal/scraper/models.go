package scraper

import "mimir-scrapper/src/internal/scraper/parser"

type Result struct {
	Index      int
	Transcript *parser.Transcript
	Error      error
}

type Institution struct {
	Name string
}

type Occasion struct {
	Name string
}

type Session struct {
	InstitutionID int
	OccasionID    int
}