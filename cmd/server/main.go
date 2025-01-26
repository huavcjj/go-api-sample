package main

import (
	"context"
	"fmt"
	"go-api-sample/domain/entity"
	"go-api-sample/infrastructure/database"
	"go-api-sample/infrastructure/web"
	"go-api-sample/pkg/logger"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := database.NewDatabase(database.InstanceMySQL)
	if err != nil {
		logger.Fatal(err.Error())
	}

	defer func() {
		if err := database.Close(db); err != nil {
			logger.Error(err.Error())
		}
	}()

	if err := database.Migrate(db, entity.NewDomains()...); err != nil {
		logger.Fatal(err.Error())
	}

	server, err := web.NewServer(web.InstanceEcho, db)
	if err != nil {
		logger.Fatal(err.Error())
	}

	go func() {
		if err := server.Start(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	<-ctx.Done()
}
