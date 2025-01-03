-- name: InsertSession :one

INSERT INTO session (id, institution_id, occasion_id, date_time)
VALUES (uuid_generate_v4(), $1, $2, $3)
RETURNING *;

-- name: FindSessionByID :one
SELECT * FROM session WHERE id = $1;