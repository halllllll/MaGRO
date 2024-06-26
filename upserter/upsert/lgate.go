package upsert

import (
	"context"
	"fmt"

	"github.com/cheggaaa/pb/v3"
	"github.com/halllllll/MaGRO/kajiki/model"
	edu "github.com/halllllll/MaGRO/kajiki/model/edu"
	"github.com/jackc/pgx/v5"
)

func (u *Upsert) LgateUpsert(ctx context.Context, target []*edu.LGateCSVOutput) error {
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	unitmap := make(map[string]int)
	bar := pb.StartNew(len(target))
	for _, v := range target {
		unit := v.SchoolName
		var unit_id int
		if _, ok := unitmap[unit]; !ok {
			// unit upsert
			err := tx.QueryRow(ctx, `INSERT INTO unit(name) VALUES($1) ON CONFLICT (name) DO UPDATE SET name = $1 RETURNING id`, unit).Scan(&unit_id)
			if err != nil && err != pgx.ErrNoRows {
				return err
			}
			unitmap[unit] = unit_id
		} else {
			unit_id = unitmap[unit]
		}

		var subunits_ids []int
		subunits := v.RowSubunits()
		for _, subunit := range subunits {
			var subunit_id int
			// subunit upsert
			if err = tx.QueryRow(ctx, `INSERT INTO subunit(unit_id, name) VALUES($1, $2) ON CONFLICT (unit_id, name) DO UPDATE SET name = $2 RETURNING id`, unit_id, subunit).Scan(&subunit_id); err != nil && err != pgx.ErrNoRows {
				return err
			}
			subunits_ids = append(subunits_ids, subunit_id)
		}

		// role

		var role model.MaGRORoleName
		// var role edu.LGateRole
		if v.IsStudent() {
			role = model.RoleManager
		} else {
			role = model.RoleGuest
		}
		var role_id int

		if err := tx.QueryRow(ctx, `SELECT id FROM role WHERE name = $1`, string(role)).Scan(&role_id); err != nil && err != pgx.ErrNoRows {
			return err
		}
		if role_id == 0 {
			return fmt.Errorf("not found role id")
		}

		var status_id int

		if err := tx.QueryRow(ctx, `SELECT id FROM status WHERE name = $1`, model.StatusActive).Scan(&status_id); err != nil && err != pgx.ErrNoRows {
			return err
		}

		// user
		if err := tx.QueryRow(ctx, `INSERT INTO users(id, account_id, name, kana, role, status) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO NOTHING`,
			v.Uuid,
			v.Username,
			fmt.Sprintf("%s %s", v.FamilyName, v.GivenName),
			fmt.Sprintf("%s %s", v.FamilyKanaName, v.GivenKanaName),
			role_id,
			status_id,
		).Scan(); err != nil && err != pgx.ErrNoRows {
			return err
		}

		// middle tables
		_, err := tx.Exec(ctx, `INSERT INTO users_unit(user_id, unit_id) VALUES($1, $2) ON CONFLICT (user_id, unit_id) DO NOTHING`, v.Uuid, unit_id)
		if err != nil && err != pgx.ErrNoRows {
			return err
		}

		// batch := pgx.Batch{}
		// for _, suid := range subunits_ids {
		// 	batch.Queue(`INSERT INTO users_subunit(user_id, subunit_id) VALUES($1, $2) ON CONFLICT (user_id, subunit_id) DO NOTHING`, v.Uuid, suid)
		// }
		// result, err := tx.SendBatch(ctx, &batch).Exec()
		// if err != nil {
		// 	return err
		// }

		for _, suid := range subunits_ids {
			_, err := tx.Exec(ctx, `INSERT INTO users_subunit(user_id, subunit_id) VALUES($1, $2) ON CONFLICT (user_id, subunit_id) DO NOTHING`, v.Uuid, suid)
			if err != nil {
				return err
			}
		}
		bar.Increment()
	}
	bar.Finish()

	return tx.Commit(ctx)
}
