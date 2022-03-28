package database

import (
	"fmt"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type PlayerModel struct {
	DB *sqlx.DB
}

func (m *PlayerModel) Insert(newPlayer models.Player) (int, error) {

	stmt := `INSERT INTO players (player_name, player_email, player_password, player_score, active, connected)
	VALUES (:player_id, :player_name, :player_score, :active, :connected)
	ON CONFLICT(player_name) DO UPDATE SET
	player_id = :player_id,
	player_name = :player_name,
	player_score = :player_score,
	active = :active,
	connected = :connected`

	// PlayerID       int32     `json:"player_id"`
	// PlayerName     string    `json:"player_name"`
	// PlayerEmail    string    `json:"player_email"`
	// PlayerPassword string    `json:"player_password"`
	// PlayerScore    int64     `json:"player_score"`
	// Active         bool      `json:"active"`
	// Connected      bool      `json:"connected"`
	// Created        time.Time `json:"created"`

	fmt.Println("insert runs...")
	fmt.Println(newPlayer)

	result, err := m.DB.NamedExec(stmt, newPlayer)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

	// _, err = db.NamedExec(`
	// 	INSERT INTO players (player_id, player_name, player_score, active, connected)
	// 	VALUES (:player_id, :player_name, :player_score, :active, :connected)
	// 	ON CONFLICT(player_name) DO UPDATE SET
	// 	player_id = :player_id,
	// 	player_name = :player_name,
	// 	player_score = :player_score,
	// 	active = :active,
	// 	connected = :connected
	// 	`, newPlayer)
	// if err != nil {
	// 	return err
	// }
	// return nil

}
