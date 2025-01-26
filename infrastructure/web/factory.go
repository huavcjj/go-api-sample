package web

import (
	"context"

	"gorm.io/gorm"
)

const (
	InstanceGin int = iota
	InstanceEcho
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

func NewServer(instance int, db *gorm.DB) (Server, error) {
	switch instance {
	case InstanceGin:
		config := NewConfigGin()
		return NewGinServer(config.Host, config.Port, config.CorsAllowOrigins, db)
	case InstanceEcho:
		config := NewConfigEcho()
		return NewEchoServer(config.Host, config.Port, config.CorsAllowOrigins, db)
	default:
		panic("invalid server instance")
	}
}
