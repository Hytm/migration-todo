package todos_repo

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

const (
	dbURL = "DB_URL"
)

var Client *pgx.Conn

func init() {
	if err := godotenv.Load("db/.env"); err != nil {
		log.Fatal(".env not found")
	}

	config, err := pgx.ParseConfig(os.Getenv(dbURL))
	if err != nil {
		log.Fatal(err)
	}

	Client, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database ready to accept connections")
}
