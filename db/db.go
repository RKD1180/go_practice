package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not open database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// deleteEventsTable()
	createTables()
}

func createTables() {
	createEventTable := `
    CREATE TABLE IF NOT EXISTS events(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date_time DATETIME NOT NULL,
        user_id INTEGER
    )
    `
	_, err := DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create event table")
	}
}

// Function to delete the events table
func deleteEventsTable() {
	_, err := DB.Exec("DROP TABLE IF EXISTS events")
	if err != nil {
		panic("Could not delete events table")
	}
}
