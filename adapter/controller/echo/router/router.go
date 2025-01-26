package router

import (
	"go-api-sample/adapter/controller/echo/handler"
	"go-api-sample/application/service"
	"go-api-sample/infrastructure/repository"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func NewEchoRouter(db *gorm.DB, corsAllowOrigins []string) (*echo.Echo, error) {
	router := echo.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	registerUserHandler(router, userHandler)

	return router, nil
}

func registerUserHandler(router *echo.Echo, userHandler *handler.UserHandler) {

}
