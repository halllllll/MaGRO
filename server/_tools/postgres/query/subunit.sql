-- name: AddSubunit :one
INSERT INTO subunit(
  unit_id, name, public
) VALUES(
  $1, $2, $3
) RETURNING *;

-- name: SearchSubunitsByUnitId :many
SELECT * FROM subunit AS s
WHERE unit_id = $1
ORDER BY s.name DESC;