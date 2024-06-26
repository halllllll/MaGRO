package upsert

import (
	"context"
	"fmt"
	"io"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Upsert struct {
	db   *pgxpool.Pool
	file io.Reader
}

func NewUpsert(db *pgxpool.Pool, file io.Reader) *Upsert {
	return &Upsert{db: db, file: file} // pb: pb.New(0),

}

// pgx動作確認
func (u *Upsert) Hoge(ctx context.Context) error {
	// bulk upsertで複数の値をreturningできる？
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	// 1: batchを使う
	//  -> 最後にRETURNINGされたものしか取れない？
	batch := &pgx.Batch{}
	for i := range 10 {
		batch.Queue(`INSERT INTO system(version) VALUES($1) RETURNING id`, fmt.Sprintf("バッチインサートのテスト %d", i))
	}

	result, err := tx.SendBatch(ctx, batch).Query()
	if err != nil {
		return err
	}
	defer result.Close()

	// TODO: なぜか最後の結果しか取得できない
	for result.Next() {
		var a int
		if err := result.Scan(&a); err != nil {
			return err
		}
	}
	if err := result.Err(); err != nil {
		return err
	}

	// 2:

	return nil
}
