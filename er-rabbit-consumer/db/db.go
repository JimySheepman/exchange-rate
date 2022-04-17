package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "merlins"
	password      = "root"
	database_name = "er"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)
	// Connect to database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return db, nil
}
