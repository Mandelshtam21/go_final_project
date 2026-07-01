package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '', 
    title VARCHAR(128) NOT NULL, 
    comment TEXT NOT NULL DEFAULT '', 
    repeat VARCHAR(128) NOT NULL DEFAULT '');

CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);`

func Init(dbFile string) error {
	dbFile = "scheduler.db"
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println("Error opening db: ", err)
		return err
	}
	return nil
	defer db.Close()
}
