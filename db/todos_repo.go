package todos_repo

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	dbURL = "DB_URL"
)

var (
	Client *pgx.Conn
	url    = os.Getenv(dbURL)
)

func init() {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	Client, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database ready to accept connections")
}
