-- name: InsertEvent :one
INSERT INTO event (id, link, record_id)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: FindEventsByRecordID :many
SELECT * FROM event WHERE record_id = $1;

-- name: FindEventByID :one
SELECT * FROM event WHERE id = $1;
