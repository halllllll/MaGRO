package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/halllllll/MaGRO/kajiki/config"
	service_reader "github.com/halllllll/MaGRO/kajiki/reader"
	"github.com/halllllll/MaGRO/kajiki/store"
	"github.com/halllllll/MaGRO/kajiki/upsert"
)

func main() {
	var dsn string
	var csvfilepath string

	csvfilepath = "./files/data.csv"

	f, err := os.Open(csvfilepath)
	if err != nil {
		log.Println("failed to open csv file. if using docker container, should bind or mount ex")
		log.Fatal(err)
	}
	defer f.Close()
	// config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("can't load env vars")
		log.Fatal(err)
	}
	// ex) 'postgres://<user>:<password>@<host>:<port>/<database name>'
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.DBUser, cfg.DBPassword, "db", cfg.DBPort, cfg.DBName)

	ctx := context.Background()

	pool, close, err := store.NewDB(ctx, dsn)
	if err != nil {
		log.Println("failed to connect db")
		log.Fatal(err)
	}
	defer close()

	upsert := upsert.NewUpsert(pool, f)
	// // pgx動作確認
	// if err := upsert.Hoge(ctx); err != nil {
	// 	log.Fatal(err)
	// }
	format := service_reader.Format(strings.ToLower(string(cfg.Service)))
	reader := service_reader.NewReader(upsert)
	start := time.Now()
	switch format {
	case service_reader.FormatLGate:
		if err := reader.Lgate(ctx, f); err != nil {
			log.Fatal(err)
		}
	default:
		panic("unsupported service")
	}
	end := time.Now()
	millisconde := end.Sub(start).Milliseconds()
	log.Printf("done: time - %d (msec)", millisconde)
}
