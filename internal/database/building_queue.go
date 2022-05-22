package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type BuildingQueueModel struct {
	DB        *sqlx.DB
	Resources []models.Resources
	Buildings []models.Buildings
}

const status10 uint8 = 10
const status20 uint8 = 20

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

// Main function
func (m *BuildingQueueModel) StartConstructionNewBuilding(buildingQueue models.BuildingQueue) error {

	building := m.getBuildingData(buildingQueue.BuildingID)

	resCheck, err := m.checkIfSufficientResources(buildingQueue, building)
	if err != nil {
		log.Println("Error checking if sufficient resources: ", err)
		return err
	}

	if resCheck {
		m.insertToBuildingQueue(buildingQueue, building)
	}
	return err
}

//
//TODO Create new function to keep track of the progress of the building queue
//
func (m *BuildingQueueModel) UpdateBuildingQueue() ([]models.BuildingRowAndVillage, error) {

	// set status of the building in the queue to start = 10 or 20

	err := m.setBuildingToStart()
	if err != nil {
		log.Println("Error setting building to start 'UpdateBuildingQueue': ", err)
		return nil, err
	}

	// check if the finish_time is reached
	// if reached, get the rowid and village_id to add the building to the village (update the village_setup table)

	rowVillageAmount, err := m.getRowAndVillageIDs()
	if err != nil {
		log.Println("Error getting row and village IDs 'UpdateBuildingQueue': ", err)
		return nil, err
	}

	return rowVillageAmount, err
}

func (m *BuildingQueueModel) getRowAndVillageIDs() ([]models.BuildingRowAndVillage, error) {

	var data []models.BuildingRowAndVillage

	err := m.DB.Select(&data, "SELECT rowid, building_id, village_id, amount FROM building_queue WHERE (status = ? OR status = ?) AND strftime('%s', 'now') >= finish_time;", status10, status20)

	if err == sql.ErrNoRows {
		err = nil
	}

	return data, err
}

func (m *BuildingQueueModel) SetBuildingToDone() error {

	fmt.Println("setBuildingToDone")

	stmt := "UPDATE building_queue SET status = status * 10 WHERE (status = ? OR status = ?) AND strftime('%s', 'now') >= finish_time;"

	_, err := m.DB.Exec(stmt, status10, status20)
	if err != nil {
		log.Println("Error setting building to done 'setBuildingToDone': ", err)
		return err
	}

	return nil
}

func (m *BuildingQueueModel) setBuildingToStart() error {

	const stmt string = "UPDATE building_queue SET status = status * 10 WHERE (status = 1 OR status = 2) AND start_time <= strftime('%s', 'now');"

	_, err := m.DB.Exec(stmt)
	if err != nil {
		log.Println("Error setting building to start 'setBuildingToStart': ", err)
		return err
	}

	return nil
}

func (m *BuildingQueueModel) getBuildingData(buildingID string) models.Buildings {

	// var building models.BuildingSQL

	for _, b := range m.Buildings {
		if b.BuildingID == buildingID {
			return b
		}
	}

	// stmt := fmt.Sprintf("SELECT * FROM buildings WHERE building_id='%s';", buildingID)

	// err := m.DB.Get(&building, stmt)
	// if err != nil {
	// 	return building, err
	// }

	return models.Buildings{}
}

func (m *BuildingQueueModel) checkIfSufficientResources(buildingQueue models.BuildingQueue, building models.Buildings) (bool, error) {

	//check if the village has sufficient resources

	// bcs := splitCostString(building.BuildCost)

	for _, bc := range building.BuildCost {

		// get the amount of the resource
		// get the resourceName
		resName, err := m.getResourceName(bc.ResourceID)
		if err != nil {
			log.Println("Error getting resource name: ", err)
			return false, err
		}

		resCheck, err := m.checkResourceFromVillage(resName, bc.Amount, buildingQueue.VillageID)
		if err != nil {
			log.Println("Error checking resource from village: ", err)
			return false, err
		}

		if resCheck {
			resUpdated, err := m.updateVillageResources(resName, bc.Amount, buildingQueue.VillageID)
			if err != nil {
				return false, err
			}
			if resUpdated {
				log.Printf("Updated village resources! ResourceName: %s, Cost: %d, VillageID: %d", resName, bc.Amount, buildingQueue.VillageID)

			}
		}
	}
	return true, nil
}

func (m *BuildingQueueModel) insertToBuildingQueue(buildingQueue models.BuildingQueue, building models.Buildings) (bool, error) {

	//TODO insert to building queue
	//buildingID, villageID, playerID, amount, status, startTime, finishTime

	const insertBuildingQueue string = `INSERT INTO building_queue (building_id, village_id, player_id, amount, status, start_time, finish_time) VALUES (:building_id, :village_id, :player_id, :amount, :status, :start_time, :finish_time)`

	buildingQueue.StartTime = uint32(time.Now().Unix())
	buildingQueue.FinishTime = buildingQueue.StartTime + buildingQueue.Amount*building.BuildTime

	tx := m.DB.MustBegin()
	_, err := tx.NamedExec(insertBuildingQueue, &buildingQueue)
	switch err {
	case nil:
		tx.Commit()
		log.Printf("Inserted building queue! BuildingID: %s, VillageID: %d, PlayerID: %d, Amount: %d, Status: %d, StartTime: %d, FinishTime: %d", buildingQueue.BuildingID, buildingQueue.VillageID, buildingQueue.PlayerID, buildingQueue.Amount, buildingQueue.Status, buildingQueue.StartTime, buildingQueue.FinishTime)
		return true, err
	default:
		log.Printf("Error inserting to building queue: %s, BuildingID: %s, VillageID: %d, PlayerID: %d, Amount: %d, Status: %d, StartTime: %d, FinishTime: %d", err, buildingQueue.BuildingID, buildingQueue.VillageID, buildingQueue.PlayerID, buildingQueue.Amount, buildingQueue.Status, buildingQueue.StartTime, buildingQueue.FinishTime)
		tx.Rollback()
		return false, err
	}
}

func (m *BuildingQueueModel) updateVillageResources(resName string, resCost uint32, villageID uint32) (bool, error) {

	stmt := fmt.Sprintf("UPDATE village_resources SET %s = %s - %d WHERE village_id=%d;", resName, resName, resCost, villageID)

	tx := m.DB.MustBegin()
	_, err := tx.Exec(stmt)
	switch err {
	case nil:
		tx.Commit()
		return true, err
	default:
		log.Printf("Error! Rollback updating village resources: %s - resName: %s, resCost: %d, villageID: %d", err, resName, resCost, villageID)
		tx.Rollback()
		return false, err
	}
}

func (m *BuildingQueueModel) checkResourceFromVillage(resourceName string, resCost uint32, villageID uint32) (bool, error) {

	// returns 1 if the village has enough resources and 0 if no sufficient resources

	var resCheck uint8
	stmt := fmt.Sprintf("SELECT CASE WHEN %s >= %d THEN '1' ELSE '0' END FROM village_resources WHERE village_id=%d;", resourceName, resCost, villageID)

	err := m.DB.Get(&resCheck, stmt)
	if err != nil {
		log.Println("Error getting resource amount: ", err)
		return false, err
	}

	if resCheck == 1 {
		return true, nil
	}

	return false, nil
}

func (m *BuildingQueueModel) getResourceName(resourceID uint32) (string, error) {

	for _, r := range m.Resources {
		if r.ResourceID == resourceID {
			return r.Resource, nil
		}
	}

	return "", nil
}

func splitCostString(s string) []models.BuildingCost {

	var bcs []models.BuildingCost

	s1 := strings.Split(s, ",")

	for _, v := range s1 {
		if v != "" {

			res := matchResource(v)
			if res == 4294967295 {
				log.Println("Error: Resource not in range")
			}
			amount := matchAmount(v)
			if amount == 4294967295 {
				log.Println("Error: Amount not in range")
			}

			bcs = append(bcs, models.BuildingCost{ResourceID: res, Amount: amount})
		}
	}
	return bcs
}

func matchResource(s string) uint32 {

	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			c := s[i+1 : j]
			c64, err := strconv.ParseUint(c, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			return uint32(c64)
		}
	}
	return 4294967295
}

func matchAmount(s string) uint32 {

	i := strings.Index(s, "[")
	if i >= 0 {
		j := strings.Index(s, "]")
		if j >= 0 {
			c := s[i+1 : j]
			c64, err := strconv.ParseUint(c, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			return uint32(c64)
		}
	}
	return 4294967295
}
