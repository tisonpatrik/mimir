// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: event.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const findEventByID = `-- name: FindEventByID :one
SELECT id, link, record_id FROM event WHERE id = $1
`

func (q *Queries) FindEventByID(ctx context.Context, id uuid.UUID) (Event, error) {
	row := q.db.QueryRow(ctx, findEventByID, id)
	var i Event
	err := row.Scan(&i.ID, &i.Link, &i.RecordID)
	return i, err
}

const findEventsByRecordID = `-- name: FindEventsByRecordID :many
SELECT id, link, record_id FROM event WHERE record_id = $1
`

func (q *Queries) FindEventsByRecordID(ctx context.Context, recordID uuid.UUID) ([]Event, error) {
	rows, err := q.db.Query(ctx, findEventsByRecordID, recordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(&i.ID, &i.Link, &i.RecordID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertEvent = `-- name: InsertEvent :one
INSERT INTO event (id, link, record_id)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING id, link, record_id
`

type InsertEventParams struct {
	Link     string    `json:"link"`
	RecordID uuid.UUID `json:"record_id"`
}

func (q *Queries) InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, insertEvent, arg.Link, arg.RecordID)
	var i Event
	err := row.Scan(&i.ID, &i.Link, &i.RecordID)
	return i, err
}
