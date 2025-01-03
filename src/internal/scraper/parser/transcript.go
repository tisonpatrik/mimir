package parser

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
