-- name: GetAction :one
SELECT * FROM action
WHERE id = $1;

-- name: CreateAction :one
INSERT INTO action(name) VALUES($1) RETURNING *;

-- name: UpdateAction :one
UPDATE action
SET name = $1
WHERE id = $2
RETURNING *;