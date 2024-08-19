package adapter

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/tshiba06/account_backend/internal/logger"
)

func NewDB() (*sqlx.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	dbSetup := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)
	ctx := context.Background()
	logger.Info(ctx, "db setup", []any{dbSetup})

	db, err := sqlx.Open("postgres", dbSetup)

	return db, err
}
