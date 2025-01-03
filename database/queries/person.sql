-- name: InsertPerson :one

INSERT INTO person (id, full_name)
VALUES (uuid_generate_v4(), $1)
RETURNING *;

-- name: FindPersonByName :one
SELECT * FROM person WHERE full_name = $1;