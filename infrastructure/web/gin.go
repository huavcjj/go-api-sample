package web

import (
	"context"
	"fmt"
	"go-api-sample/adapter/controller/gin/router"

	"net/http"

	"gorm.io/gorm"
)

type ginServer struct {
	server *http.Server
}

func (g *ginServer) Start() error {
	return g.server.ListenAndServe()
}

func (g *ginServer) Shutdown(ctx context.Context) error {
	return g.server.Shutdown(ctx)
}

func NewGinServer(host, port string, corsAllowOrigins []string, db *gorm.DB) (Server, error) {
	ginRouter, err := router.NewGinRouter(db, corsAllowOrigins)
	if err != nil {
		return nil, err
	}
	return &ginServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: ginRouter,
		},
	}, nil
}
