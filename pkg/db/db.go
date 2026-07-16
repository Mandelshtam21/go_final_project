package db

import (
	"database/sql"
	"fmt"
	"path/filepath"

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
	var err error

	absPath, err := filepath.Abs(dbFile)
	if err != nil {
		return err
	}

	fmt.Println("SERVER DATABASE:", absPath)

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println("Error opening db: ", err)
		return err
	}

	_, err = db.Exec(schema)
	if err != nil {
		fmt.Println("Error creating db schema: ", err)
		return err
	}

	return nil
}
