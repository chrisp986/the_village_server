package database

import (
	"fmt"
	"log"
	"time"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type VillageSetupModel struct {
	DB *sqlx.DB
}

func (m *VillageSetupModel) Insert(newVillageSetup models.VillageSetup) (uint32, error) {

	stmt := `INSERT OR IGNORE INTO village_setup (village_id, player_id, buildings, status, last_update) VALUES (:village_id, :player_id, :buildings, :status, :last_update)`

	result, err := m.DB.NamedExec(stmt, &newVillageSetup)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (m *VillageSetupModel) InsertWithIDCheck(village_id uint32, player_id uint32) (uint32, error) {

	newVillageSetup := models.VillageSetup{
		VillageID:  village_id,
		PlayerID:   player_id,
		Buildings:  "", //TODO fix this
		Status:     0,
		LastUpdate: time.Now().Local().String(),
	}

	stmt := `INSERT OR IGNORE INTO village_setup (village_id, player_id, buildings, status, last_update) VALUES (:village_id, :player_id, :buildings, :status, :last_update)`

	result, err := m.DB.NamedExec(stmt, &newVillageSetup)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (m *VillageSetupModel) Update(newVillageSetup models.VillageSetup) (uint32, error) {

	stmt := `INSERT OR IGNORE INTO village_setup (village_id, player_id, buildings, status, last_update) VALUES (:village_id, :player_id, :buildings, :status, :last_update)`

	result, err := m.DB.NamedExec(stmt, &newVillageSetup)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func SetBuildings() string {

	//TODO fix this

	// "0|0,1|0,2|0,3|0,4|0,5|0,6|0,7|0,8|0,9|0,10|0,11|0,12|0,13|0,14|0,15|0,16|0,17|0,18|0,19|0,20|0,21|0,22|0,23|0,24|0"

	var new string
	// var count int
	// m.DB.Get(&count, "SELECT COUNT(*) FROM buildings")

	for i := 0; i < 25; i++ {
		new += fmt.Sprintf("%d|0,", i)
	}

	return new
}
