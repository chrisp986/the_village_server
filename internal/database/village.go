package database

import (
	"log"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type VillageModel struct {
	DB *sqlx.DB
}

func (m *VillageModel) Insert(newVillage models.Village) (int, error) {

	stmt := `INSERT OR IGNORE INTO villages (player_id, village_name, village_size, village_status, village_loc_y, village_loc_x) VALUES (:player_id, :village_name, :village_size, :village_status, :village_loc_y, :village_loc_x);`

	result, err := m.DB.NamedExec(stmt, &newVillage)
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
