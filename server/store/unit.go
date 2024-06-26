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

func (r *Repository) ListUsersSubunits(ctx context.Context, unitid *entity.UnitId) ([]db.GetUsersSubunitsRow, error) {
	uid := int32(*unitid)
	result, err := r.query.GetUsersSubunits(ctx, uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}
