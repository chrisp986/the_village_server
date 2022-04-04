package database

import (
	"log"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type PlayerResourcesModel struct {
	DB *sqlx.DB
}

func (m *PlayerResourcesModel) Insert(newPlayerResources models.PlayerResource) (int, error) {

	// CREATE TABLE IF NOT EXISTS player_resources (
	// 	village_id INTEGER PRIMARY KEY,
	// 	player_id INTEGER NOT NULL,
	// 	food INTEGER NOT NULL,
	// 	wood INTEGER NOT NULL,
	// 	stone INTEGER NOT NULL,
	// 	copper INTEGER NOT NULL,
	// 	water INTEGER NOT NULL,
	// 	gold INTEGER NOT NULL,
	// 	UNIQUE(village_id)
	// 	);

	stmt := `INSERT OR IGNORE INTO player_resources (village_id, player_id, food, wood, stone, copper, water, gold) VALUES (:village_id, :player_id, :food, :wood, :stone, :copper, :water, :gold);`

	result, err := m.DB.NamedExec(stmt, &newPlayerResources)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
