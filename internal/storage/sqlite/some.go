package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER
    );
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	return db
}
