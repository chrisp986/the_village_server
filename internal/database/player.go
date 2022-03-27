package database

import (
	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type PlayerModel struct {
	DB *sqlx.DB
}

func (m *PlayerModel) Insert(newPlayer models.Player) (int, error) {

	stmt := `INSERT INTO players (player_id, player_name, player_score, active, connected)
	VALUES (:player_id, :player_name, :player_score, :active, :connected)
	ON CONFLICT(player_name) DO UPDATE SET
	player_id = :player_id,
	player_name = :player_name,
	player_score = :player_score,
	active = :active,
	connected = :connected`

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
