-- name: InsertOccasion :one

INSERT INTO occasion (id, name)
VALUES (uuid_generate_v4(), $1)
RETURNING *;

-- name: FindOccasionByName :one
SELECT * FROM occasion WHERE name = $1;