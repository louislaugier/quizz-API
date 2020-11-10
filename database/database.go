package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Postgres
)

// DB export
var DB = connect()

func connect() *sql.DB {
	// if err := godotenv.Load(); err != nil {
	// 	panic(err)
	// }
	db, _ := sql.Open("postgres", "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":5432/"+os.Getenv("DB_NAME"))
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
