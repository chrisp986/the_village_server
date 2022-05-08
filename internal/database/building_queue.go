package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

	//check if the village has sufficient resources

	sufficientResources, resourceNames, err := m.checkIfSufficientResources(buildingQueue, building)
	if err != nil {
		log.Println("Error checking if sufficient resources: ", err)
		return err
	}
	fmt.Println(sufficientResources)
	fmt.Println(resourceNames)
	return err
}

func (m *BuildingQueueModel) getBuildingData(buildingID string) (models.BuildingSQL, error) {

	var building models.BuildingSQL

	stmt := fmt.Sprintf("SELECT * FROM buildings WHERE building_id='%s';", buildingID)

	err := m.DB.Get(&building, stmt)
	if err != nil {
		return building, err
	}

	return building, err
}

func (m *BuildingQueueModel) checkIfSufficientResources(buildingQueue models.BuildingQueue, building models.BuildingSQL) (bool, []string, error) {

	//check if the village has sufficient resources

	var resourceName []string
	bcs := splitCostString(building.BuildCost)

	for _, bc := range bcs {
		fmt.Println("ResourceID: ", bc.ResourceID)
		fmt.Println("Amount: ", bc.Amount)
		// get the amount of the resource
		// get the resourceName
		resName, err := m.getResourceName(bc.ResourceID)
		if err != nil {
			log.Println("Error getting resource name: ", err)
			return false, nil, err
		}

		// check if the village has enough resources
		resourceName = append(resourceName, resName)
	}
	return true, resourceName, nil
}

func (m *BuildingQueueModel) checkResourceFromVillage(resourceName string, villageID uint32) (bool, error) {

	// returns 1 if the village has enough resources and 0 if no sufficient resources

	var resAmount uint8
	stmt := fmt.Sprintf("SELECT CASE WHEN %s >= 20 THEN '1' ELSE '0' END FROM village_resources WHERE village_id=%d;", resourceName, villageID)

	err := m.DB.Get(&resAmount, stmt)
	if err != nil {
		log.Println("Error getting resource amount: ", err)
		return false, err
	}

	if resAmount == 1 {
		return true, nil
	}

	return false, nil
}

func (m *BuildingQueueModel) getResourceName(resourceID uint32) (string, error) {

	var resourceName string
	stmt := fmt.Sprintf("SELECT resource FROM resources WHERE resource_id='%d';", resourceID)

	err := m.DB.Get(&resourceName, stmt)
	if err != nil {
		return "", err
	}
	return resourceName, nil
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
