-- name: GetStatusName :one
SELECT * FROM status
WHERE id = $1;

-- name: CreateStatusName :one
INSERT INTO status(name) VALUES($1) RETURNING *;

-- name: UpdateStatusName :one
UPDATE status
SET name = $1
WHERE id = $2
RETURNING *;