package store

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
	db "github.com/halllllll/MaGRO/gen/sqlc"
	"github.com/halllllll/MaGRO/store/transaction"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetSystemInfo(ctx context.Context) (*entity.System, error) {
	info := &entity.System{}
	result, err := r.query.GetSystemInfo(ctx)
	if err != nil {
		return nil, err
	}
	info.Version = entity.SystemVersion(result.Version)
	info.Created = result.CreatedAt
	info.Modified = result.UpdatedAt

	return info, nil
}

// admin用なやつ あれば追加していく
func (r *Repository) UpdateRole(ctx context.Context, req *entity.ReqNewRoleAlias) error {
	tr := transaction.NewTransaction(r.pool)
	// いいやり方がわかんないからわかんないことはやらないことにした
	a := string(req.AliasAdminName)
	d := string(req.AliasDirectorName)
	m := string(req.AliasManagerName)
	g := string(req.AliasGuestName)
	var args []*db.UpdateRoleNameParams
	if a != "" {
		args = append(args, &db.UpdateRoleNameParams{
			Name:      "admin",
			NameAlias: &req.AliasAdminName,
		})
	}
	if d != "" {
		args = append(args, &db.UpdateRoleNameParams{
			Name:      "director",
			NameAlias: &req.AliasDirectorName,
		})
	}
	if m != "" {
		args = append(args, &db.UpdateRoleNameParams{
			Name:      "manager",
			NameAlias: &req.AliasManagerName,
		})
	}
	if g != "" {
		args = append(args, &db.UpdateRoleNameParams{
			Name:      "guest",
			NameAlias: &req.AliasGuestName,
		})
	}

	err := tr.WithTransaction(ctx, func(tx pgx.Tx) error {
		q := r.query.WithTx(tx)
		for _, arg := range args {
			_, err := q.UpdateRoleName(ctx, *arg)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err

}
