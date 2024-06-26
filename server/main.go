package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/config"
)

//go:embed static/*
var static embed.FS

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %+v", err)
	}

}

type Ping struct {
	Status  int       `json:"status"`
	Cur     time.Time `json:"timestamp"`
	Message string    `json:"message"`
}

func hello(ctx *gin.Context) {
	p := Ping{Status: http.StatusOK, Cur: time.Now(), Message: "hello!!!"}
	ctx.JSON(http.StatusOK, p)
	return
}

func htmlHello(ctx *gin.Context) {
	w := ctx.Writer
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
	return
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port: %d: %+v", cfg.Port, err)
	}
	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {
		return err
	}
	s := NewServer(l, mux)
	return s.Run(ctx)

}
