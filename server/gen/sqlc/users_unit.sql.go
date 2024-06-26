// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users_unit.sql

package db

import (
	"context"
)

const mapUserUnit = `-- name: MapUserUnit :exec
INSERT INTO users_unit(
  user_id, unit_id
) VALUES($1, $2)
`

type MapUserUnitParams struct {
	UserID string `db:"user_id" json:"user_id"`
	UnitID int32  `db:"unit_id" json:"unit_id"`
}

func (q *Queries) MapUserUnit(ctx context.Context, arg MapUserUnitParams) error {
	_, err := q.db.Exec(ctx, mapUserUnit, arg.UserID, arg.UnitID)
	return err
}
