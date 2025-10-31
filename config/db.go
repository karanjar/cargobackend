package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDb() {

	dsn := "user=dustin password=12345 dbname=cardb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error openning database")
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	Db = db
}
