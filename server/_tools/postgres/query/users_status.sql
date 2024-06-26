-- name: GetStatusName :one
SELECT * FROM users_status
WHERE id = $1;

-- name: CreateStatusName :one
INSERT INTO users_status(name) VALUES($1) RETURNING *;

-- name: UpdateStatusName :one
UPDATE users_status
SET name = $1
WHERE id = $2
RETURNING *;