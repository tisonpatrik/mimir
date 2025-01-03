-- name: InsertInstitution :one

INSERT INTO institution (id, name)
VALUES (uuid_generate_v4(), $1)
RETURNING *;