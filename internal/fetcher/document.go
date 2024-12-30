package fetcher

type Document struct {
	Period           string       // Období
	Session          string       // Schůze
	Committee        string       // Orgán
	PrintNumber      string       // Č. tisku
	ResolutionNumber string       // Č. usnesení
	Title            string       // Dokument
	Date             string       // Datum
	DocumentType     DocumentType // Typ dokumentu
	ApproxSize       string       // Přibližná velikost
	Preview          string       // Náhled
	Original         string       // Originál
}

type DocumentType struct {
	Name   string
	Format FileFormat
}

type FileFormat string

const (
	FormatPDF     FileFormat = "PDF"
	FormatHTML    FileFormat = "HTML"
	FormatUnknown FileFormat = "unknown" // For unknown or new formats
)

const (
	DocumentTypeSeminar    = "seminář"  // Seminar
	DocumentTypeInvitation = "pozvánka" // Invitation
	DocumentTypeResolution = "usnesení" // Resolution
	DocumentTypeTranscript = "steno"    // Transcript
	DocumentTypeMinutes    = "zápis"    // Minutes
	DocumentTypeUnknown    = "unknown"  // Generic type for unknown document types
)
