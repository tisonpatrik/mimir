package parser

import "github.com/google/uuid"

type Record struct {
	SessionID      uuid.UUID `json:"session_id"`
	SpeakerName    string    `json:"string"`
	Content        string    `json:"content"`
	Events         []Event   `json:"events"`
	SequenceNumber int       `json:"sequence_number"`
}

type Event struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Speaker struct {
	Name string `json:"name"`
}
