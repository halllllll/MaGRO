-- name: UpdateApp :one
UPDATE app
SET (
  title, client_id, unit_alias, subunit_alias
) = ($1, $2, $3, $4)
WHERE id = 1 -- 常に一行目に期待されるデータが格納されているとする
RETURNING *;

-- name: CreateAppInfo :one
INSERT INTO app(
  title, client_id, unit_alias, subunit_alias
)VALUES($1, $2, $3, $4)
RETURNING *;