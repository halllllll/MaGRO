package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	service_reader "github.com/halllllll/MaGRO/kajiki/reader"
	"github.com/halllllll/MaGRO/kajiki/store"
	"github.com/halllllll/MaGRO/kajiki/upsert"
)

func main() {
	var dsn string
	var csvfilepath string
	var service string
	flag.StringVar(&dsn, "dsn", "", "connect to postgresql database. ex) 'postgres://<user>:<password>@<host>:<port>/<database name>'")
	flag.StringVar(&service, "service", "", "target service: 'lgate', 'loilo', 'c4th', 'miraiseed'")
	flag.StringVar(&csvfilepath, "csv", "", "csv file path based on Format")
	flag.Parse()

	f, err := os.Open(csvfilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ctx := context.Background()

	pool, close, err := store.NewDB(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	upsert := upsert.NewUpsert(pool, f)
	// // pgx動作確認
	// if err := upsert.Hoge(ctx); err != nil {
	// 	log.Fatal(err)
	// }
	format := service_reader.Format(strings.ToLower(service))
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
	fmt.Printf("done: time - %d (msec)", millisconde)
}
