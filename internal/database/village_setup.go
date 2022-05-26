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

type VillageSetupModel struct {
	DB      *sqlx.DB
	Workers []models.Workers
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

	buildingID := m.getBuildingsID()

	newVillageSetup := models.VillageSetup{
		VillageID:  village_id,
		PlayerID:   player_id,
		WorkerID:   InitBuildingsString(buildingID),
		Status:     0,
		LastUpdate: time.Now().Local().Format("2006-01-02 15:04:05"),
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

func (m *VillageSetupModel) getBuildingsID() []string {

	var wID []string

	for _, b := range m.Workers {
		wID = append(wID, b.WorkerID)
	}
	return wID
}

func InitBuildingsString(buildingID []string) string {

	var new string

	for _, v := range buildingID {
		new += fmt.Sprintf("%s=0,", v)
	}

	return new

	//TODO work on this to make it dynamic e.g. if a new building is added to the game, it will be added to the database

	// buildingsString, err := getBuildingsString(db)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 	}

	// 	log.Println("Error while getting building string.", err)
	// 	return ""
	// }

	// splitBuildingsString(buildingsString)

	// bs := strings.Split(buildingsString, ",")

	// if len(bs)-1 == buildingsCount {
	// 	fmt.Println("Buildings string is already initialized.", len(bs), buildingsCount)

	// } else {
	// 	fmt.Println("Building string is not initialized.")
	// }

	// buildingsCount, err := getBuildingsCount(db)
	// if err != nil {
	// 	log.Println("Error while getting building count.", err)
	// 	return ""
	// }

}

func (m *VillageSetupModel) GetBuildingCount(villageID uint32) (string, error) {

	var bString string
	err := m.DB.Get(&bString, "SELECT buildings FROM village_setup WHERE village_id = ?;", villageID)
	if err != nil {
		log.Printf("Error while getting building string: %v ;villageID: %d", err, villageID)
		return "", err
	}

	return bString, err
}

func (m *VillageSetupModel) UpdateBuildingString(bString string, brv models.BuildingRowAndVillage) (bool, error) {

	// bcs := splitString(bString)

	// for _, v := range bcs {
	// 	// fmt.Println("Building:", v.BuildingID, " Count:", v.Count)
	// 	if v.BuildingID == brv.BuildingID {
	// 		// fmt.Println("Building:", v.BuildingID, " Count:", v.Count)
	// 		newBString := addCountToBuilding(bString, brv.BuildingID, brv.Amount)
	// 		// fmt.Println("New Building String:", newBString)
	// 		err := m.updateBuildingStringWithNewAmount(newBString, brv.VillageID)
	// 		if err != nil {
	// 			log.Printf("Error while updating building string: %v ;villageID: %d", err, brv.VillageID)
	// 			return false, err
	// 		}
	// 	}

	// }
	return true, nil
}

func (m *VillageSetupModel) updateBuildingStringWithNewAmount(bString string, villageID uint32) error {

	stmt := "UPDATE village_setup SET buildings = '?' WHERE village_id = ?;"

	tx := m.DB.MustBegin()
	_, err := tx.Exec(stmt, bString, villageID)
	switch err {
	case nil:
		tx.Commit()
		log.Printf("Successfully updated building string for the village: %d", villageID)
		return err
	default:
		log.Printf("Error while updating building string for the village: %d", villageID)
		tx.Rollback()
		return err
	}

}

func addCountToBuilding(buildingString string, buildingID string, count uint32) string {

	s1 := strings.Index(buildingString, buildingID+"=")

	if s1 >= 0 {

		s2 := strings.Index(buildingString[s1+3:], ",")

		if s2 >= 0 {

			sn := buildingString[s1+3:]

			sn1 := sn[:s2]

			b64, err := strconv.ParseUint(sn1, 10, 32)
			if err != nil {
				log.Fatal(err)
			}

			newCount := fmt.Sprintf("%d", count+uint32(b64))

			newString := fmt.Sprintf("%s%s%s", buildingString[:s1+3], newCount, sn[s2:])

			return newString
		}
	}
	log.Fatalln("Error in addCountToBuilding")
	return buildingString

}

// func VerifyBuildingsString(db *sqlx.DB) error {

// 	buildingsString, err := getBuildingsString(db)
// 	if err != nil {
// 		log.Println("Error while getting building string.", err)
// 		return err
// 	}

// 	buildingsCount, err := getBuildingsCount(db)
// 	if err != nil {
// 		log.Println("Error while getting building count.", err)
// 		return err
// 	}

// 	bs := strings.Split(buildingsString, ",")

// 	if len(bs)-1 == buildingsCount {
// 		return nil
// 	} else {
// 		// updateBuildingsString(db, buildingsString)
// 	}

// 	return errors.New("Building string is not initialized.")
// }

func (m *VillageSetupModel) getBuildings(village_id uint32) (models.BuildingCount, error) {

	var buildingCount models.BuildingCount
	stmt := fmt.Sprintf(`SELECT buildings FROM village_setup WHERE village_id = %d;`, village_id)

	err := m.DB.Get(&buildingCount, stmt)

	return buildingCount, err
}

// func (m *VillageSetupModel) updateBuildingsString(bs string) (bool, error) {

// 	return true, err
// }

// func (m *VillageSetupModel) updateBuildings(village_id uint32, building_id uint32, change uint32) {

// 	vs := models.VillageSetup{}
// 	m.DB.Get(&vs, "SELECT * FROM village_setup WHERE village_id = ?", village_id)

// 	index := strings.Index(vs.Buildings, fmt.Sprintf("(%d)", building_id))

// }
