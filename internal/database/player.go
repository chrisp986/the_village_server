package database

import (
	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type PlayerModel struct {
	DB *sqlx.DB
}

func (m *PlayerModel) Insert(newPlayer models.Player) (int32, error) {

	stmt := `INSERT INTO players (player_name, player_email, player_password, player_score, active, connected, created)
	VALUES (?, ?, ?, ?, ?, ?, datetime('now','localtime'));`

	// PlayerID       int32     `json:"player_id"`
	// PlayerName     string    `json:"player_name"`
	// PlayerEmail    string    `json:"player_email"`
	// PlayerPassword string    `json:"player_password"`
	// PlayerScore    int64     `json:"player_score"`
	// Active         bool      `json:"active"`
	// Connected      bool      `json:"connected"`
	// Created        time.Time `json:"created"`

	result, err := m.DB.Exec(stmt, newPlayer.PlayerName, newPlayer.PlayerEmail, newPlayer.PlayerPassword, newPlayer.PlayerScore, newPlayer.Active, newPlayer.Connected)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

func (m *PlayerModel) Get(pID int) (*models.Player, error) {

	p := &models.Player{}

	stmt := `SELECT player_id, player_name, player_email, player_password, player_score, active, connected, created FROM players WHERE player_id = ?;`

	// err := m.DB.QueryRowx(stmt, pID).StructScan(&p)
	err := m.DB.QueryRow(stmt, pID).Scan(&p.PlayerID, &p.PlayerName, &p.PlayerEmail, &p.PlayerPassword, &p.PlayerScore, &p.Active, &p.Connected, &p.Created)
	if err != nil {
		return nil, err
	}
	return p, err
}
