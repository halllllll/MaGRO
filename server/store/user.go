package store

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
	db "github.com/halllllll/MaGRO/gen/sqlc"
	"github.com/halllllll/MaGRO/store/transaction"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) RegisterUser(ctx context.Context, u *entity.User_T) (*entity.User_T, error) {

	// トランザクションを使う
	tr := transaction.NewTransaction(r.pool)
	args := &db.AddUserTParams{
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role,
		Created:  r.Clocker.Now(),
		Modified: r.Clocker.Now(),
	}

	var user *entity.User_T
	err := tr.WithTransaction(ctx, func(tx pgx.Tx) error {
		txQuery := r.query.WithTx(tx)
		newUser, err := txQuery.AddUserT(ctx, *args)
		if err != nil {
			return err
		}
		user = &entity.User_T{
			ID:       entity.UserID(newUser.ID),
			Name:     newUser.Name,
			Role:     newUser.Role,
			Created:  newUser.Created,
			Modified: newUser.Modified,
		}
		return nil
	})
	return user, err
}
