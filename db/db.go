package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalDB *sql.DB

func InitDB() {
	conn, err := sql.Open("sqlite3", "api.db")

	GlobalDB = conn

	if err != nil {
		panic("could not connect to the database")
	}

	fmt.Println("successfully connected to the database")

	GlobalDB.SetMaxOpenConns(10)
	GlobalDB.SetMaxIdleConns(5)

	// create EVENTS table
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := GlobalDB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}
}
