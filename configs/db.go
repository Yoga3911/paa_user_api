package configs

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func DatabaseConnection() *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}

	dsn := os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Println(err.Error())
	}

	config.MaxConns = 20

	pg, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println(err.Error())
	}

	action := 1
	switch action {
	case 1:
		log.Println("Migration")
		migration(pg)
	case 2:
		log.Println("Rollback")
		rollback(pg)
	}

	return pg
}
