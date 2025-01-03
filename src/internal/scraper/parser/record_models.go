package parser

type Record struct {
	SpeakerName    string  `json:"string"`
	Content        string  `json:"content"`
	Events         []Event `json:"events"`
	SequenceNumber int     `json:"sequence_number"`
}

type Event struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
