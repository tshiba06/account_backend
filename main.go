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
	"github.com/tshiba06/account_backend/internal/telemetry"
	roleRepo "github.com/tshiba06/account_backend/repository/role"
	userRepo "github.com/tshiba06/account_backend/repository/user"
	roleUC "github.com/tshiba06/account_backend/usecase/role"
	userUC "github.com/tshiba06/account_backend/usecase/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

func main() {
	tp, err := telemetry.New()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	db, err := adapter.NewDB()
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}

	tracer := otel.Tracer("gin-server")

	// repository
	roleRepo := roleRepo.NewRepository(db)
	userRepo := userRepo.NewRepository(db, tracer)
	fmt.Println("roleRepository", roleRepo)

	// usecase
	roleUseCase := roleUC.NewUseCase(roleRepo)
	userUseCase := userUC.NewUsecase(tracer, userRepo)
	fmt.Println("roleUseCase", roleUseCase)

	router := gin.Default()
	router.ContextWithFallback = true
	router.Use(otelgin.Middleware("account-server"))

	handlers := &adapter.Handler{
		Tracer:      tracer,
		UserUseCase: userUseCase,
	}

	api.RegisterHandlers(router, handlers)

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
