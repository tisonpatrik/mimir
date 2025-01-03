package services

import (
	"context"
	"encoding/json"
	"log"
	"mimir-scrapper/src/internal/scraper/parser"
	"mimir-scrapper/src/pkg/repository"

	"github.com/google/uuid"
)

type RecordsService struct {
	Queries *repository.Queries
}

// NewRecordsService initializes a new RecordsService with the given Queries instance.
func NewRecordsService(queries *repository.Queries) *RecordsService {
	return &RecordsService{
		Queries: queries,
	}
}

// ProcessAndSaveRecord processes and saves a single record and its events.
func (rs *RecordsService) ProcessAndSaveRecord(ctx context.Context, sessionID uuid.UUID, record parser.Record) error {
	// Find or create speaker
	person, err := rs.Queries.FindPersonByName(ctx, record.SpeakerName)
	if err != nil {
		if err.Error() == "no rows in result set" {
			log.Printf("Speaker '%s' not found, creating it.", record.SpeakerName)
			person, err = rs.Queries.InsertPerson(ctx, record.SpeakerName)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Serialize events to JSON
	eventsJSON, err := json.Marshal(record.Events)
	if err != nil {
		log.Printf("Error serializing events for record: %v", err)
		return err
	}

	// Create record
	newRecord, err := rs.Queries.InsertRecord(ctx, repository.InsertRecordParams{
		SessionID:      sessionID,
		SpeakerID:      person.ID,
		Content:        record.Content,
		Events:         eventsJSON,
		SequenceNumber: int32(record.SequenceNumber),
	})
	if err != nil {
		log.Printf("Error inserting record: %v", err)
		return err
	}

	// Save events
	for _, event := range record.Events {
		_, err := rs.Queries.InsertEvent(ctx, repository.InsertEventParams{
			Link:     event.Link,
			RecordID: newRecord.ID, // Directly use newRecord.ID as uuid.UUID
		})
		if err != nil {
			log.Printf("Error inserting event for record ID %s: %v", newRecord.ID, err)
			return err
		}
	}

	return nil
}
