package store

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
	db "github.com/halllllll/MaGRO/gen/sqlc"
)

func (r *Repository) ListUnits(ctx context.Context, id *entity.UserID) ([]db.Unit, error) {
	uid := string(*id)
	units, err := r.query.GetBelongsUnits(ctx, uid)
	if err != nil {
		return nil, err
	}
	return units, nil
}

func (r *Repository) ListUsersSubunits(ctx context.Context, userUuid *entity.UserUUID, unitId *entity.UnitId) ([]db.GetSubunitsByUserUuIDAndUnitIdRow, error) {
	unit_id := int32(*unitId)
	user_uuid := string(*userUuid)

	args := &db.GetSubunitsByUserUuIDAndUnitIdParams{
		ID:     unit_id,   // unit id
		UserID: user_uuid, // user uuid
	}

	result, err := r.query.GetSubunitsByUserUuIDAndUnitId(ctx, *args)
	// TOOD: pgx.ErrNoRows handling
	if err != nil {
		return nil, err
	}
	return result, nil
}
