package store

import (
	"context"
	"fmt"
	"time"

	db "github.com/halllllll/MaGRO/gen/sqlc"

	"github.com/halllllll/MaGRO/config"

	"github.com/halllllll/MaGRO/clock"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, cfg *config.Config) (db.Querier, func(), error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.DBUser, cfg.DBPassword, "db", cfg.DBPort, cfg.DBName)
	pgxConf, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, func() { pool.Close() }, err
	}
	xdb := db.New(pool)

	return xdb, func() { pool.Close() }, nil
}

type Repository struct {
	Clocker clock.RealClocker
	pool    *pgxpool.Pool
	query   *db.Queries
}

// sqlcで生成したものに依存するんじゃなくてpgxpoolにしたほうがいいんじゃないかというやつ
func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, func(), error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.DBUser, cfg.DBPassword, "db", cfg.DBPort, cfg.DBName)
	pgxConf, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, func() { pool.Close() }, err
	}
	return pool, func() { pool.Close() }, nil
}

// DB作るやつ..
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Clocker: clock.RealClocker{},
		pool:    pool,
		query:   db.New(pool),
	}
}
