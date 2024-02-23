package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=root password=tejljdfoau8uojlkj dbname=example sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(db)


	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
