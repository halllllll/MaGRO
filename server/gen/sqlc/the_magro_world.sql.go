// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: the_magro_world.sql

package db

import (
	"context"
	"time"
)

const aggregateUnitInfo = `-- name: AggregateUnitInfo :many
SELECT un.name AS "unit", COUNT(DISTINCT su.name) AS "subunit_count", COUNT(DISTINCT u_s.user_id) AS "user_count"
FROM unit AS un
JOIN subunit AS su ON su.unit_id = un.id
JOIN users_subunit AS u_s ON u_s.subunit_id = su.id
GROUP BY un.id ORDER BY un.name
`

type AggregateUnitInfoRow struct {
	Unit         string `db:"unit" json:"unit"`
	SubunitCount int64  `db:"subunit_count" json:"subunit_count"`
	UserCount    int64  `db:"user_count" json:"user_count"`
}

// for admin
func (q *Queries) AggregateUnitInfo(ctx context.Context) ([]AggregateUnitInfoRow, error) {
	rows, err := q.db.Query(ctx, aggregateUnitInfo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AggregateUnitInfoRow{}
	for rows.Next() {
		var i AggregateUnitInfoRow
		if err := rows.Scan(&i.Unit, &i.SubunitCount, &i.UserCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBelongsUnits = `-- name: GetBelongsUnits :many
SELECT un.id, un.name FROM unit AS un
JOIN users_unit AS uu ON un.id = uu.unit_id
JOIN users AS u ON uu.user_id = u.id
WHERE u.account_id = $1
`

// by user id
func (q *Queries) GetBelongsUnits(ctx context.Context, accountID string) ([]Unit, error) {
	rows, err := q.db.Query(ctx, getBelongsUnits, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Unit{}
	for rows.Next() {
		var i Unit
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubunitByUserID = `-- name: GetSubunitByUserID :many


SELECT unit.name AS "unit", su.name AS "subunit", su.id AS "subunit_id",su.public AS "public", su.created_at, su.updated_at, users.id AS "user_id", users.account_id, users.name AS "user_name", users.kana AS "user_kananame", role.name AS "role", COALESCE(NULLIF(role.name_alias, ''), role.name) AS "role_alias"
FROM (
    SELECT id, unit_id, name, public, created_at, updated_at FROM subunit WHERE unit_id = $1
) AS su
INNER JOIN users_subunit AS us ON su.id = us.subunit_id
INNER JOIN users ON users.id = us.user_id
INNER JOIN unit ON su.unit_id = unit.id
INNER JOIN role ON users.role = role.id
ORDER BY unit.name ASC, su.name ASC
`

type GetSubunitByUserIDRow struct {
	Unit         string    `db:"unit" json:"unit"`
	Subunit      string    `db:"subunit" json:"subunit"`
	SubunitID    int32     `db:"subunit_id" json:"subunit_id"`
	Public       bool      `db:"public" json:"public"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UserID       string    `db:"user_id" json:"user_id"`
	AccountID    string    `db:"account_id" json:"account_id"`
	UserName     string    `db:"user_name" json:"user_name"`
	UserKananame *string   `db:"user_kananame" json:"user_kananame"`
	Role         string    `db:"role" json:"role"`
	RoleAlias    string    `db:"role_alias" json:"role_alias"`
}

// `u.account_id` for all account
// by unit id
func (q *Queries) GetSubunitByUserID(ctx context.Context, unitID int32) ([]GetSubunitByUserIDRow, error) {
	rows, err := q.db.Query(ctx, getSubunitByUserID, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubunitByUserIDRow{}
	for rows.Next() {
		var i GetSubunitByUserIDRow
		if err := rows.Scan(
			&i.Unit,
			&i.Subunit,
			&i.SubunitID,
			&i.Public,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.AccountID,
			&i.UserName,
			&i.UserKananame,
			&i.Role,
			&i.RoleAlias,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubunitsByUserUuIDAndUnitId = `-- name: GetSubunitsByUserUuIDAndUnitId :many
SELECT un.name AS "unit", su.name AS "subunit", su.id AS "subunit_id",su.public AS "public", su.created_at, su.updated_at, users.id AS "user_id", users.account_id, users.name AS "user_name", users.kana AS "user_kananame", role.name AS "role", COALESCE(NULLIF(role.name_alias, ''), role.name) AS "role_alias"
FROM users
JOIN users_subunit AS u_s ON users.id = u_s.user_id
JOIN subunit AS su ON u_s.subunit_id = su.id
JOIN unit AS un ON un.id = su.unit_id
JOIN role ON users.role = role.id
WHERE su.id IN (
  SELECT su2.id
  FROM subunit AS su2
  JOIN users_subunit AS u_s2 ON su2.id = u_s2.subunit_id
  WHERE u_s2.user_id = $1
)
AND un.id = $2
ORDER BY un.name ASC, su.name ASC
`

type GetSubunitsByUserUuIDAndUnitIdParams struct {
	UserID string `db:"user_id" json:"user_id"`
	ID     int32  `db:"id" json:"id"`
}

type GetSubunitsByUserUuIDAndUnitIdRow struct {
	Unit         string    `db:"unit" json:"unit"`
	Subunit      string    `db:"subunit" json:"subunit"`
	SubunitID    int32     `db:"subunit_id" json:"subunit_id"`
	Public       bool      `db:"public" json:"public"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UserID       string    `db:"user_id" json:"user_id"`
	AccountID    string    `db:"account_id" json:"account_id"`
	UserName     string    `db:"user_name" json:"user_name"`
	UserKananame *string   `db:"user_kananame" json:"user_kananame"`
	Role         string    `db:"role" json:"role"`
	RoleAlias    string    `db:"role_alias" json:"role_alias"`
}

// AND users.id != $1 自分は含むことにする
func (q *Queries) GetSubunitsByUserUuIDAndUnitId(ctx context.Context, arg GetSubunitsByUserUuIDAndUnitIdParams) ([]GetSubunitsByUserUuIDAndUnitIdRow, error) {
	rows, err := q.db.Query(ctx, getSubunitsByUserUuIDAndUnitId, arg.UserID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubunitsByUserUuIDAndUnitIdRow{}
	for rows.Next() {
		var i GetSubunitsByUserUuIDAndUnitIdRow
		if err := rows.Scan(
			&i.Unit,
			&i.Subunit,
			&i.SubunitID,
			&i.Public,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.AccountID,
			&i.UserName,
			&i.UserKananame,
			&i.Role,
			&i.RoleAlias,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
