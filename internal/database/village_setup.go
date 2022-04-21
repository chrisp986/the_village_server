package database

import (
	"errors"
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

func GetBuildingsID(db *sqlx.DB) ([]string, error) {

	var bID []string
	err := db.Select(&bID, "SELECT building_id FROM buildings ORDER BY rowid;")

	return bID, err
}

// func splitBuildingsString(bs string) []models.BuildingCount {

// 	var bcs []models.BuildingCount
// 	s := strings.Split(bs, ",")

// 	for _, v := range s {
// 		s1 := strings.Split(v, ")")
// 		b := strings.Replace(s1[0], "(", "", -1)
// 		if b == "" {
// 			continue
// 		}
// 		b64, err := strconv.ParseUint(b, 10, 32)
// 		if err != nil {
// 			log.Fatal(err)

// 		}
// 		replacer := strings.NewReplacer("[", "", "]", "")
// 		c := replacer.Replace(s1[1])
// 		if c == "" {
// 			continue
// 		}
// 		c64, err := strconv.ParseUint(c, 10, 32)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		bc := models.BuildingCount{
// 			BuildingID: uint32(b64),
// 			Count:      uint32(c64),
// 		}
// 		bcs = append(bcs, bc)

// 	}

// 	return bcs
// }

func splitBuildingsString(s string) []models.BuildingCount {

	var bcs []models.BuildingCount

	s1 := strings.Split(s, ",")

	for _, v := range s1 {
		if v != "" {

			b := matchB(v)
			if b == "" {
				log.Println("Error: Building ID not in range")
			}
			c := matchC(v)
			if c == 4294967295 {
				log.Println("Error: Count not in range")
			}

			bcs = append(bcs, models.BuildingCount{
				BuildingID: b,
				Count:      c,
			})
		}
	}
	return bcs
}

func matchB(s string) string {

	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			b := s[i+1 : j]
			return b
		}
	}
	return ""
}

func matchC(s string) uint32 {

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

func InitBuildingsString(db *sqlx.DB, bID []string) string {

	var new string

	for _, v := range bID {
		new += fmt.Sprintf("(%s)[0],", v)

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

func VerifyBuildingsString(db *sqlx.DB) error {

	buildingsString, err := getBuildingsString(db)
	if err != nil {
		log.Println("Error while getting building string.", err)
		return err
	}

	buildingsCount, err := getBuildingsCount(db)
	if err != nil {
		log.Println("Error while getting building count.", err)
		return err
	}

	bs := strings.Split(buildingsString, ",")

	if len(bs)-1 == buildingsCount {
		return nil
	} else {
		updateBuildingsString(db, buildingsString)
	}

	return errors.New("Building string is not initialized.")
}

func updateBuildingsString(db *sqlx.DB, bs string) error {

	return nil
}

// func (m *VillageSetupModel) updateBuildings(village_id uint32, building_id uint32, change uint32) {

// 	vs := models.VillageSetup{}
// 	m.DB.Get(&vs, "SELECT * FROM village_setup WHERE village_id = ?", village_id)

// 	index := strings.Index(vs.Buildings, fmt.Sprintf("(%d)", building_id))

// }
