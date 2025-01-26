package web

import (
	"context"
	"fmt"
	"go-api-sample/adapter/controller/echo/router"
	"go-api-sample/pkg/logger"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type echoServer struct {
	router *echo.Echo
}

func (e *echoServer) Start() error {
	return e.router.Start(e.router.Server.Addr)
}

func (e *echoServer) Shutdown(ctx context.Context) error {
	return e.router.Shutdown(ctx)
}

func NewEchoServer(host, port string, corsAllowOrigins []string, db *gorm.DB) (Server, error) {
	echoRouter, err := router.NewEchoRouter(db, corsAllowOrigins)
	if err != nil {
		logger.Error(err.Error(), "host", host, "port", port)
		return nil, err
	}
	echoRouter.Server.Addr = fmt.Sprintf("%s:%s", host, port)
	return &echoServer{
		router: echoRouter,
	}, nil
}
