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

func getBuildingsCount(db *sqlx.DB) (int, error) {

	var bc int
	err := db.Get(&bc, "SELECT COUNT(*) FROM buildings;")

	return bc, err
}

func getBuildingsString(db *sqlx.DB) (string, error) {

	var bs string
	err := db.Get(&bs, "SELECT buildings FROM village_setup LIMIT 1;")

	return bs, err
}

func splitBuildingsString(bs string) []models.BuildingCount {

	var bcs []models.BuildingCount
	s := strings.Split(bs, ",")

	for _, v := range s {
		s1 := strings.Split(v, ")")
		b := strings.Replace(s1[0], "(", "", -1)
		if b == "" {
			continue
		}
		b64, err := strconv.ParseUint(b, 10, 32)
		if err != nil {
			log.Fatal(err)

		}
		replacer := strings.NewReplacer("[", "", "]", "")
		c := replacer.Replace(s1[1])
		if c == "" {
			continue
		}
		c64, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		bc := models.BuildingCount{
			BuildingID: uint32(b64),
			Count:      uint32(c64),
		}
		bcs = append(bcs, bc)

	}

	return bcs
}

func InitBuildingsString(db *sqlx.DB) string {

	buildingsCount, err := getBuildingsCount(db)
	if err != nil {
		log.Println("Error while getting building count.", err)
		return ""
	}

	buildingsString, err := getBuildingsString(db)
	if err != nil {
		log.Println("Error while getting building string.", err)
		return ""
	}

	splitBuildingsString(buildingsString)

	// bs := strings.Split(buildingsString, ",")

	// if len(bs)-1 == buildingsCount {
	// 	fmt.Println("Buildings string is already initialized.", len(bs), buildingsCount)

	// } else {
	// 	fmt.Println("Building string is not initialized.")
	// }

	// "0|0,1|0,2|0,3|0,4|0,5|0,6|0,7|0,8|0,9|0,10|0,11|0,12|0,13|0,14|0,15|0,16|0,17|0,18|0,19|0,20|0,21|0,22|0,23|0,24|0"

	var new string
	// var count int
	// m.DB.Get(&count, "SELECT COUNT(*) FROM buildings")

	for i := 0; i < buildingsCount; i++ {
		new += fmt.Sprintf("(%d)[0],", i)
	}

	return new
}

// func (m *VillageSetupModel) updateBuildings(village_id uint32, building_id uint32, change uint32) {

// 	vs := models.VillageSetup{}
// 	m.DB.Get(&vs, "SELECT * FROM village_setup WHERE village_id = ?", village_id)

// 	index := strings.Index(vs.Buildings, fmt.Sprintf("(%d)", building_id))

// }
