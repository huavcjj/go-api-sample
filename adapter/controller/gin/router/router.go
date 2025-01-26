package router

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(db *gorm.DB, corsAllowOrigins []string) (*gin.Engine, error) {
	router := gin.Default()

	//userRepository := repository.NewUserRepository(db)
	//userService := service.NewUserService(userRepository)
	//userHandler := handler.NewUserHandler(userService)
	//
	//registerUserHandler(router, userHandler)

	return router, nil
}

//func registerUserHandler(router *gin.Engine, userHandler *handler.UserHandler) {
//
//}
