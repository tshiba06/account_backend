package main

import (
	"fmt"
	"log"

	"github.com/tshiba06/account_backend/api"
	roleRepo "github.com/tshiba06/account_backend/repository/role"
	"github.com/tshiba06/account_backend/server"
	roleUC "github.com/tshiba06/account_backend/usecase/role"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=root password=tejljdfoau8uojlkj dbname=example sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	// repository
	roleRepository := roleRepo.NewRepository(db)
	fmt.Println(roleRepository)

	// usecase
	roleUseCase := roleUC.NewUseCase(roleRepository)
	fmt.Println(roleUseCase)

	router := gin.Default()

	srv := server.Server{}

	api.RegisterHandlers(router, srv)

	// TODO: graceful shutdown

	if err := router.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
