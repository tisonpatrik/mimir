package handlers

import "mimir-scrapper/internal/parser"

type Result struct {
	Index      int
	Transcript *parser.Transcript
	Error      error
}
