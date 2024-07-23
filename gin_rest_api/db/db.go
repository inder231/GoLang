package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Variable to store the db connection
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Cound not connect to database!")
	}

	DB.SetMaxOpenConns(10) // max connection which can be open simultaneously
	DB.SetMaxIdleConns(5) // It will set 5 connection to be open ideally if no one is using these connections

	// Create Tables
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

	// Create table
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

}
