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

	gin_static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/auth"
	"github.com/halllllll/MaGRO/config"
	"github.com/halllllll/MaGRO/handler"
	"github.com/halllllll/MaGRO/service"
	"github.com/halllllll/MaGRO/store"

	// gin_static "github.com/soulteary/gin-static"

	"golang.org/x/sync/errgroup"
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

// 意味がわからないが、mainパッケージの別ファイルや別パッケージだとdocker buildでエラーになる
// ＊NewServer, NewMuxはmainパッケージの別ファイルにおいていた

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// ここは別ゴルーチン
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("faield to shutdown: %+v", err)
	}
	return eg.Wait()

}

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// michiからginに移行
	router := gin.Default()
	router.ContextWithFallback = true
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// csrfはいったんおいておく
	// （SPAなのでサーバー側でinputやmetaタグにトークンをぶちこんだりすることができない）
	// https://zenn.dev/leaner_dev/articles/20210930-rails-api-spa-csrf
	// https://kimuson.dev/blog/%E3%83%95%E3%83%AD%E3%83%B3%E3%83%88%E3%82%A8%E3%83%B3%E3%83%89/csrf_spa/
	// むしろ不要な気がしている
	dbPool, cleanup, err := store.NewPool(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	repo := store.NewRepository(dbPool)
	at := handler.AddTask{
		Service: &service.AddTask{
			Repo: repo,
		},
	}

	router.Use(gin_static.Serve("/", gin_static.EmbedFolder(static, "static")))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/redirect", func(ctx *gin.Context) {
		fmt.Println("👹　ここはredirect")
		fmt.Printf("request👺: %#v\n", ctx.Request)
		// fmt.Printf("uri: %s\n", ctx.Request.RequestURI)
		// ctx.Request.URL.Path = "/"
		// router.HandleContext(ctx)
		ctx.Redirect(http.StatusFound, "/")
	})

	router.GET("/login", func(ctx *gin.Context) {
		fmt.Println("🍊　ここはlogin")
		fmt.Printf("request: %#v\n", ctx.Request)
		// fmt.Printf("uri: %s\n", ctx.Request.RequestURI)
		// ctx.Request.URL.Path = "/"
		// router.HandleContext(ctx)
		ctx.Redirect(http.StatusOK, "/")
	})

	router.POST("/tasks", at.AddTask)

	lt := handler.ListTask{
		Service: &service.ListTask{
			Repo: repo,
		},
	}
	router.GET("/tasks", lt.ListTask)

	// MaGRO
	ms := handler.MagroSystem{
		Service: &service.MagroSystem{
			Repo: repo,
		},
	}

	// 非ゲストアカウント用
	// Middlewareでわけなくてもいい気がする(guestでやらせたいことをいい感じに分けたいが今のところguestはなにもできない)
	ensureRegularAccount := auth.NewEnsureRegularAccountMiddleware(repo)

	magro := router.Group("/api").Use(auth.MsalAuthMiddleware(cfg.ClientId))

	magro.GET("/info", ms.GetSystemInfoHandler)

	mu := handler.MaGROUnitList{
		Service: &service.ListUnit{
			Repo: repo,
		},
	}

	// TODO: repass用	仮実装
	mr := handler.MaGRORepass{
		Service: &service.Repass{
			Repo: repo,
		},
	}

	maguroRegular := magro.Use(ensureRegularAccount.EnsureRegularAccountMiddleWare())

	maguroRegular.GET("/units", mu.ListUnit)
	maguroRegular.GET("/subunit/:unit", mu.ListUsersSubunit)

	maguroRegular.POST("/units/:unit/repass", mr.Repass)

	ma := handler.MaGROAdmin{
		MutateService: &service.MutateMAGRO{
			Repo: repo,
		},
	}

	maguroRegular.PUT("/role/new", ma.UpdateRoleNameHandler)

	return router, cleanup, err
}
