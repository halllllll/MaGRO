package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: sqlcを使わないpgx/v5で愚直にxlsxからデータをDBにupsertするよ

func NewDB(ctx context.Context, dsn string) (*pgxpool.Pool, func(), error) {
	// conn, err := pgxpool.New(ctx, dsn)
	pgxConf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, nil, err

	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, func() { pool.Close() }, err

	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, func() { pool.Close() }, err
	}

	return pool, func() { pool.Close() }, err
}

// // sqlcで生成したものに依存するんじゃなくてpgxpoolにしたほうがいいんじゃないかというやつ
// func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, func(), error) {
// 	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.DBUser, cfg.DBPassword, "db", cfg.DBPort, cfg.DBName)
// 	pgxConf, err := pgxpool.ParseConfig(uri)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	defer cancel()
// 	if err := pool.Ping(ctx); err != nil {
// 		return nil, func() { pool.Close() }, err
// 	}
// 	return pool, func() { pool.Close() }, nil
// }

// // DB作るやつ..
// func NewRepository(pool *pgxpool.Pool) *Repository {
// 	return &Repository{
// 		Clocker: clock.RealClocker{},
// 		pool:    pool,
// 		query:   db.New(pool),
// 	}
// }
