-- name: AddUser :one
INSERT INTO users(
  id, account_id, name, kana, role, status
) VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET (
  account_id, name, kana, role, status
) = ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByRole :many
SELECT * FROM users AS u
INNER JOIN role AS r ON r.id = u.role
WHERE r.name = $1;

-- name: GetUserByStatus :many
SELECT * FROM users AS u
INNER JOIN users_status AS us ON us.id = u.status
WHERE us.name = $1;

-- name: SearchUserByAccountID :many
SELECT * FROM users AS u
WHERE LOWER(u.account_id) LIKE '%' || $1 || '%'
ORDER BY u.account_id DESC;

