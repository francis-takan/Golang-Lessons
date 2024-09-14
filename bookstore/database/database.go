package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DataBase() *sql.DB {
	//use Enviroment variables
	connstr := os.Getenv("DATABASE_URL")
	if len(connstr) == 0 {
		log.Fatal("Failed to load enviroment variables")
	}

	db, err := sql.Open("postgres", connstr)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}
