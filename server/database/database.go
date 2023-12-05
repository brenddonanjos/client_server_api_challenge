package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func StartConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/sqlite.db")
	if err != nil {
		return db, err
	}

	criarTabela := `CREATE TABLE IF NOT EXISTS exchange_rate_history (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			conversion TEXT NOT NULL,
			bid DECIMAL(13,2),
			created_at DATETIME NOT NULL
		);`

	_, err = db.Exec(criarTabela)
	if err != nil {
		db.Close()
		return db, err
	}

	return db, nil
}
