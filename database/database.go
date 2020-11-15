package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres
)

// DB export
var DB = connect()

func connect() *sql.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
