package main

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const db_file string = "tv_server.db"

const createTable string = `
  CREATE TABLE IF NOT EXISTS players (
  player_id INTEGER NOT NULL PRIMARY KEY,
  player_name TEXT NOT NULL,
  player_score INTEGER NOT NULL,
  active INTEGER NOT NULL,
  connected INTEGER NOT NULL,
  UNIQUE(player_name)
  );`

func initDB() error {
	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(createTable); err != nil {
		return err
	}
	return nil
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", db_file)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}
	return db, nil
}
