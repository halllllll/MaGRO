package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/auth"
	"github.com/halllllll/MaGRO/config"
	"github.com/halllllll/MaGRO/handler"
	"github.com/halllllll/MaGRO/service"
	"github.com/halllllll/MaGRO/store"
)

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

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
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

	magro := router.Group("/api").Use(auth.MsalAuthMiddleware(cfg.ClientId))

	magro.GET("/info", ms.GetSystemInfoHandler)

	mu := handler.MaGROUnitList{
		Service: &service.ListUnit{
			Repo: repo,
		},
	}
	magro.GET("/units", mu.ListUnit)
	magro.GET("/subunits", mu.ListUsersSubunit)

	ma := handler.MaGROAdmin{
		MutateService: &service.MutateMAGRO{
			Repo: repo,
		},
	}

	magro.PUT("/role/new", ma.UpdateRoleNameHandler)

	return router, cleanup, err
}
