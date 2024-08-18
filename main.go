package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tshiba06/account_backend/adapter"
	"github.com/tshiba06/account_backend/api"
	roleRepo "github.com/tshiba06/account_backend/repository/role"
	roleUC "github.com/tshiba06/account_backend/usecase/role"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := adapter.NewDB()
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}
	defer db.Close()

	// repository
	roleRepository := roleRepo.NewRepository(db)
	fmt.Println("roleRepository", roleRepository)

	// usecase
	roleUseCase := roleUC.NewUseCase(roleRepository)
	fmt.Println("roleUseCase", roleUseCase)

	router := gin.Default()

	api.RegisterHandlers(router, &adapter.Handler{})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
