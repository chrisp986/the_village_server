package database

import (
	"log"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type VillageSetupModel struct {
	DB *sqlx.DB
}

func (m *VillageSetupModel) Insert(newVillageSetup models.VillageSetup) (int, error) {

	stmt := `INSERT OR IGNORE INTO village_setup (village_id, player_id, hunterhut_1, hunterhut_2, hunterhut_3, hunterhut_4, hunterhut_5, woodcutterhut_1, woodcutterhut_2, woodcutterhut_3, woodcutterhut_4, woodcutterhut_5,quarry_1, quarry_2, quarry_3, quarry_4, quarry_5, coppermine_1, coppermine_2, coppermine_3, coppermine_4, coppermine_5, fountain_1, fountain_2, fountain_3, fountain_4, fountain_5) VALUES (:village_id, :player_id, :hunterhut_1, :hunterhut_2, :hunterhut_3, :hunterhut_4, :hunterhut_5, :woodcutterhut_1, :woodcutterhut_2, :woodcutterhut_3, :woodcutterhut_4, :woodcutterhut_5, :quarry_1, :quarry_2, :quarry_3, :quarry_4, :quarry_5, :coppermine_1, :coppermine_2, :coppermine_3, :coppermine_4, :coppermine_5, :fountain_1, :fountain_2, :fountain_3, :fountain_4, :fountain_5);`

	result, err := m.DB.NamedExec(stmt, &newVillageSetup)
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
