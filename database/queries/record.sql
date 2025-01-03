-- name: InsertRecord :one
INSERT INTO record (id, session_id, speaker_id, content, events, sequence_number)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)
RETURNING *;

-- name: FindRecordByID :one
SELECT * FROM record WHERE id = $1;

-- name: FindRecordsBySessionID :many
SELECT * FROM record WHERE session_id = $1 ORDER BY sequence_number;

-- name: FindRecordsBySpeakerID :many
SELECT * FROM record WHERE speaker_id = $1 ORDER BY sequence_number;