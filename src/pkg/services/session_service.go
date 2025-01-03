package services

import (
	"context"
	"log"
	"mimir-scrapper/src/pkg/repository"
	"time"
)

type SessionService struct {
	Queries *repository.Queries
}

// NewSessionService initializes a new SessionServices with the given Queries instance.
func NewSessionService(queries *repository.Queries) *SessionService {
	return &SessionService{
		Queries: queries,
	}
}

// GetOrCreateSession retrieves or creates the necessary session, institution, and occasion.
func (ss *SessionService) GetOrCreateSession(ctx context.Context, institutionName, occasionName string, session_time time.Time) (*repository.Session, error) {
	// Find or create institution
	institution, err := ss.Queries.FindInstitutionByName(ctx, institutionName)
	if err != nil {
		if err.Error() == "no rows in result set" {
			log.Printf("Institution '%s' not found, creating it.", institutionName)
			institution, err = ss.Queries.InsertInstitution(ctx, institutionName)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Find or create occasion
	occasion, err := ss.Queries.FindOccasionByName(ctx, occasionName)
	if err != nil {
		if err.Error() == "no rows in result set" {
			log.Printf("Occasion '%s' not found, creating it.", occasionName)
			occasion, err = ss.Queries.InsertOccasion(ctx, occasionName)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Create session
	session, err := ss.Queries.InsertSession(ctx, repository.InsertSessionParams{
		InstitutionID: institution.ID,
		OccasionID:    occasion.ID,
		DateTime:      session_time,
	})
	if err != nil {
		return nil, err
	}

	return &session, nil
}
