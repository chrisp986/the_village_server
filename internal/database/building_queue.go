package database

import (
	"fmt"
	"log"
	"time"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type BuildingQueueModel struct {
	DB *sqlx.DB
}

// CREATE TABLE IF NOT EXISTS building_queue (
// 	building_id TEXT NOT NULL,
// 	village_id INTEGER NOT NULL,
// 	player_id INTEGER NOT NULL,
// 	amount INTEGER NOT NULL,
// 	status INTEGER NOT NULL,
// 	start_time TEXT NOT NULL,
// 	finish_time TEXT NOT NULL
// 	);

func (m *BuildingQueueModel) Insert(newBuilding models.BuildingQueue) (uint32, error) {

	stmt := `INSERT INTO building_queue (building_id, village_id, player_id, amount, status, start_time, finish_time)
	VALUES (:building_id, :village_id, :player_id, :amount, :status, :start_time, :finish_time);`

	newBuilding.StartTime = uint32(time.Now().Unix())

	result, err := m.DB.NamedExec(stmt, &newBuilding)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (m *BuildingQueueModel) StartConstructionNewBuilding(buildingQueue models.BuildingQueue) error {

	building, err := m.getBuildingData(buildingQueue.BuildingID)
	if err != nil {
		log.Println("Error getting building data: ", err)
		return err
	}
	log.Println("Building: ", building)
	return err
}

func (m *BuildingQueueModel) getBuildingData(buildingID string) (models.BuildingSQL, error) {

	var building models.BuildingSQL

	log.Println("Building ID: ", buildingID)

	stmt := fmt.Sprintf("SELECT * FROM buildings WHERE building_id='%s';", buildingID)

	err := m.DB.Get(&building, stmt)
	if err != nil {
		return building, err
	}

	return building, err
}
