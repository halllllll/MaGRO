-- name: AddUser :one
INSERT INTO users(
  id, account_id, name, kana, role, status
) VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetUserByUUID :one
SELECT users.id, users.account_id, users.name, users.kana, role.name AS "role", status.name AS "status", users.created_at, users.updated_at
FROM users
JOIN role ON role.id = users.role
JOIN status ON role.id = status.id
WHERE users.id = $1;

-- uuidから引く前提なので
-- name: ReverseLookupUUID :one
SELECT users.id
FROM users
WHERE users.account_id = $1;

-- name: UpdateUser :one
UPDATE users
SET (
  account_id, name, kana, role, status
) = ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListUsersByRole :many
SELECT * FROM users AS u
INNER JOIN role AS r ON r.id = u.role
WHERE r.name = $1;

-- name: ListUsersByStatus :many
SELECT * FROM users AS u
INNER JOIN status AS s ON s.id = u.status
WHERE s.name = $1;

-- name: SearchUserByAccountID :many
SELECT * FROM users AS u
WHERE LOWER(u.account_id) LIKE '%' || $1 || '%'
ORDER BY u.account_id DESC;


-- name: GetUserRole :one
SELECT u.account_id, u.name, r.name AS "role"
FROM users AS u JOIN role AS r ON u.role = r.id WHERE u.id = $1;