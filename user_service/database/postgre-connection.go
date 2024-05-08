package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetPostgreConnection() *sqlx.DB {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	args := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", "localhost", "5432", "root", "root", "admin", "disable")

	db, err := sqlx.Connect("postgres", args)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	return db
}
