package store

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
)

// account_idからuuidを引く
func (r *Repository) Me(ctx context.Context, id *entity.UserID) (*entity.UserUUID, error) {
	uid := string(*id)
	u, err := r.query.ReverseLookupUUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	uu := entity.UserUUID(u)
	return &uu, nil
}

func (r *Repository) GetRole(ctx context.Context, id *entity.UserID) (entity.Role, error) {
	uuid, err := r.Me(ctx, id)
	if err != nil {
		return "", err
	}
	data, err := r.query.GetUserRole(ctx, string(*uuid))
	if err != nil {
		return "", err
	}
	role := entity.Role(data.Role)
	return role, nil
}
