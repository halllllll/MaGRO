-- by user id
-- name: GetBelongsUnits :many
SELECT un.id, un.name FROM unit AS un
JOIN users_unit AS uu ON un.id = uu.unit_id
JOIN users AS u ON uu.user_id = u.id
WHERE u.account_id = $1; -- `u.account_id` for all account


-- by unit id
-- name: GetUsersSubunits :many
SELECT unit.name AS "unit", su.name AS "subunit", su.id AS "subunit_id",su.public AS "public", su.created_at, su.updated_at, users.id AS "user_id", users.account_id, users.name AS "user_name", users.kana AS "user_kananame", role.name AS "role", COALESCE(NULLIF(role.name_alias, ''), role.name) AS "role_alias"
FROM (
    SELECT * FROM subunit WHERE unit_id = $1
) AS su
INNER JOIN users_subunit AS us ON su.id = us.subunit_id
INNER JOIN users ON users.id = us.user_id
INNER JOIN unit ON su.unit_id = unit.id
INNER JOIN role ON users.role = role.id
ORDER BY unit.name ASC, su.name ASC;


-- for admin
-- name: AggregateUnitInfo :many
SELECT un.name AS "unit", COUNT(DISTINCT su.name) AS "subunit_count", COUNT(DISTINCT u_s.user_id) AS "user_count"
FROM unit AS un
JOIN subunit AS su ON su.unit_id = un.id
JOIN users_subunit AS u_s ON u_s.subunit_id = su.id
GROUP BY un.id ORDER BY un.name;