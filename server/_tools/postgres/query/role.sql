-- name: AddRole :one
INSERT INTO role(name) VALUES($1)
RETURNING *;

-- name: GetRole :one
SELECT * FROM role
WHERE id = $1;


-- name: UpdateRoleName :one
UPDATE role
SET name_alias = COALESCE($1, name_alias)
WHERE name = $2
RETURNING *;

-- name: UserRoleSet :many
SELECT * FROM role;