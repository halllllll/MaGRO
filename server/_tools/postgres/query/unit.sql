-- name: AddUnit :one
INSERT INTO unit(
  name
) VALUES ($1)
RETURNING *;

-- name: GetUnit :one
SELECT * FROM unit
WHERE id = $1;
