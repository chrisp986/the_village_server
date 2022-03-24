package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *sql.DB
}

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

func InitDB() (*App, error) {
	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}
	return &App{
		DB: db,
	}, nil
}
