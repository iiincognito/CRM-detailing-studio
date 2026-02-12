package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := os.Getenv("CONN_DB")
	config, err := pgxpool.ParseConfig(dbConn)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 100
	config.MinConns = 10
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		defer pool.Close()
		return nil, err
	}
	return pool, nil
}
