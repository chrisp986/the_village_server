package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/chrisp986/the_village_server/internal/database"
	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const db_file string = "tv_server.db"

const createTable string = `
	CREATE TABLE IF NOT EXISTS genesis (
		genesis_tick INTEGER NOT NULL,
		status INTEGER NOT NULL
	);

  CREATE TABLE IF NOT EXISTS players (
  	player_id INTEGER PRIMARY KEY,
  	player_name TEXT NOT NULL,
  	player_email TEXT NOT NULL,
  	player_password TEXT NOT NULL,
  	player_score INTEGER NOT NULL,
  	active INTEGER NOT NULL,
  	connected INTEGER NOT NULL,
  	created TEXT NOT NULL,
  	UNIQUE(player_name)
  );

CREATE TABLE IF NOT EXISTS village_resources (
	village_id INTEGER PRIMARY KEY, 
	player_id INTEGER NOT NULL,
	food INTEGER NOT NULL,
	wood INTEGER NOT NULL,
	stone INTEGER NOT NULL,
	metal INTEGER NOT NULL,
	water INTEGER NOT NULL,
	gold INTEGER NOT NULL,
	UNIQUE(village_id)
	);
  
  CREATE TABLE IF NOT EXISTS villages (
	village_id INTEGER PRIMARY KEY,
	player_id INTEGER NOT NULL,
	village_name TEXT NOT NULL,
	village_size INTEGER NOT NULL,
	village_status INTEGER NOT NULL,
	village_loc_y INTEGER NOT NULL,
	village_loc_x INTEGER NOT NULL,
	UNIQUE(village_id)
  );

CREATE TABLE IF NOT EXISTS village_setup (
	village_id INTEGER NOT NULL PRIMARY KEY, 
	player_id INTEGER NOT NULL,
	buildings TEXT NOT NULL,
	status INTEGER NOT NULL,
	last_update TEXT NOT NULL,
	UNIQUE(village_id)
	);


CREATE TABLE IF NOT EXISTS building_queue (
	building_id TEXT NOT NULL,
	village_id INTEGER NOT NULL,
	player_id INTEGER NOT NULL,
	amount INTEGER NOT NULL,
	status INTEGER NOT NULL,
	start_time TEXT NOT NULL,
	finish_time TEXT NOT NULL
	);
  `

//   CREATE TABLE IF NOT EXISTS resources (
// 	resource_id INTEGER PRIMARY KEY,
// 	resource TEXT NOT NULL,
// 	quality INTEGER NOT NULL
//   );

//   CREATE TABLE IF NOT EXISTS buildings (
// 	building_id TEXT NOT NULL,
// 	name TEXT NOT NULL,
// 	quality INTEGER NOT NULL,
// 	resource_id INTEGER NOT NULL,
// 	production_rate INTEGER NOT NULL,
// 	build_cost TEXT NOT NULL,
// 	upgrade_cost TEXT NOT NULL,
// 	build_time TEXT NOT NULL,
// 	upgrade_time TEXT NOT NULL,
// 	UNIQUE(building_id)
//   );

//   CREATE TABLE IF NOT EXISTS prod_buildings_cfg (
// 	building_id INTEGER PRIMARY KEY,
// 	resource TEXT NOT NULL,
// 	quality INTEGER NOT NULL,
// 	res_rate INTEGER NOT NULL,
// 	res_1 INTEGER NOT NULL,
// 	cost_res_1 INTEGER NOT NULL,
// 	res_2 INTEGER NOT NULL,
// 	cost_res_2 INTEGER NOT NULL,
// 	res_3 INTEGER NOT NULL,
// 	cost_res_3 INTEGER NOT NULL,
// 	res_4 INTEGER NOT NULL,
// 	cost_res_4 INTEGER NOT NULL,
// 	res_5 INTEGER NOT NULL,
// 	cost_res_5 INTEGER NOT NULL,
// 	UNIQUE(building_id)
//   );
//   CREATE TABLE IF NOT EXISTS resource_rates (
// 	resource_ INTEGER NOT NULL,
// 	quality INTEGER NOT NULL,
// 	rate INTEGER NOT NULL
//   );

const insertVillageSetup string = `INSERT OR IGNORE INTO village_setup (village_id, player_id, buildings, status, last_update) VALUES (:village_id, :player_id, :buildings, :status, :last_update)`

const insertResourceRates string = `INSERT OR IGNORE INTO resource_rates (resource, quality, rate) VALUES (:resource, :quality, :rate)`

const insertVillages string = `INSERT OR IGNORE INTO villages (village_id, player_id, village_name, village_size, village_status, village_loc_y, village_loc_x) VALUES (:village_id, :player_id, :village_name, :village_size, :village_status, :village_loc_y, :village_loc_x)`

const insertVillageSetupInit string = `INSERT OR IGNORE INTO village_setup (village_id, player_id, buildings, status, last_update) VALUES (:village_id, :player_id, :buildings, :status, :last_update)`

const insertVillageResourcesInit string = `INSERT OR IGNORE INTO village_resources (village_id, player_id, food, wood, stone, metal, water, gold) VALUES (:village_id, :player_id, :food, :wood, :stone, :metal, :water, :gold);`

// const insert string = `INSERT INTO prod_buildings_cfg (building_id,	resource, quality, res_rate,
// 	res_1,cost_res_1,res_2, cost_res_2, res_3, cost_res_3, res_4, cost_res_4, res_5, cost_res_5) VALUES (:building_id, :resource, :quality, :res_rate, :res_1, :cost_res_1, :res_2, :cost_res_2, :res_3, :cost_res_3, :res_4, :cost_res_4, :res_5, :cost_res_5);`

func initDB() error {
	db, err := sqlx.Open("sqlite3", db_file)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(createTable); err != nil {
		return err
	}
	return nil
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", db_file)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	return db, nil
}

const insertResources string = `INSERT OR IGNORE INTO resources (resource_id, resource, quality) VALUES (:resource_id, :resource, :quality)`

func initResourceTable(db *sqlx.DB) {

	res := resourcesTable()

	for _, r := range res {
		_, err := db.NamedExec(insertResources, &r)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

// const insertBuildings string = `INSERT OR IGNORE INTO buildings (building_id, name, quality, resource_id, production_rate, build_cost, upgrade_cost, build_time, upgrade_time) VALUES (:building_id, :name, :quality, :resource_id, :production_rate, :build_cost, :upgrade_cost, :build_time, :upgrade_time);`

// func initBuildingsTable(db *sqlx.DB) {

// 	buildingsJSON := buildingsTable()

// 	// log.Println(buildingsJSON)

// 	var buildings []models.BuildingSQL

// 	for _, b := range buildingsJSON {
// 		buildings = append(buildings, models.BuildingSQL{
// 			BuildingID:     b.BuildingID,
// 			Name:           b.Name,
// 			Quality:        b.Quality,
// 			ResourceID:     b.ResourceID,
// 			ProductionRate: b.ProductionRate,
// 			BuildCost:      costToString(b.BuildCost),
// 			UpgradeCost:    costToString(b.UpgradeCost),
// 			BuildTime:      b.BuildTime,
// 			UpgradeTime:    b.UpgradeTime,
// 		})

// 	}

// 	for _, b := range buildings {

// 		_, err := db.NamedExec(insertBuildings, &b)
// 		if err != nil {
// 			log.Fatalln("Error inserting building: ", err)
// 		}

// 	}
// }

func costToString(bc []models.BuildingCost) string {

	var cost string
	for _, v := range bc {
		cost += fmt.Sprintf("(%d)[%d],", v.ResourceID, v.Amount)
	}

	return cost
}

func buildingsTable() []models.Buildings {

	file := filepath.FromSlash("./internal/database/init/buildings.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var buildings []models.Buildings
	err = json.Unmarshal(bytes, &buildings)

	if err != nil {
		fmt.Println("JSON decode error in 'buildingsTable'", err)
		return nil
	}
	return buildings
}

func initVillageTable(db *sqlx.DB) {

	vil := villageTable()
	// bID, err := getBuildingsID(db)
	var bID []string
	buildings := buildingsTable()

	for _, b := range buildings {
		bID = append(bID, b.BuildingID)
	}

	for _, v := range vil {
		_, err := db.NamedExec(insertVillages, &v)
		if err != nil {
			log.Fatalln("Error inserting village: ", err)
		}

		vs := models.VillageSetup{
			VillageID:  v.VillageID,
			PlayerID:   v.PlayerID,
			Buildings:  database.InitBuildingsString(bID),
			Status:     0,
			LastUpdate: time.Now().Local().Format("2006-01-02 15:04:05"),
		}

		_, err = db.NamedExec(insertVillageSetupInit, &vs)
		if err != nil {
			log.Fatalln("Error inserting village setup", err)
		}

		vr := models.VillageResource{
			VillageID: v.VillageID,
			PlayerID:  v.PlayerID,
			Food:      100,
			Wood:      100,
			Stone:     100,
			Metal:     100,
			Water:     100,
			Gold:      20,
		}

		_, err = db.NamedExec(insertVillageResourcesInit, &vr)
		if err != nil {
			log.Fatalln("Error inserting village setup", err)
		}

	}
}

// func getBuildingsID(db *sqlx.DB) ([]string, error) {

// 	var bID []string
// 	err := db.Select(&bID, "SELECT building_id FROM buildings ORDER BY rowid;")

// 	return bID, err
// }

func resourcesTable() []models.Resources {

	file := filepath.FromSlash("./internal/database/init/resources.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var res []models.Resources
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		fmt.Println("JSON decode error!", err)
		return nil
	}
	return res
}

func villageTable() []models.Village {

	rand.Seed(time.Now().UnixNano())

	file := filepath.FromSlash("./internal/database/init/villages.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var vil []models.Village

	err = json.Unmarshal(bytes, &vil)
	if err != nil {
		fmt.Println("JSON decode error!", err)
		return nil
	}

	for i := 0; i < len(vil); i++ {
		vil[i].VillageLocY = uint32(randInt(0, 100))
		vil[i].VillageLocX = uint32(randInt(0, 100))
	}

	return vil
}

func playerTable() []models.Player {

	file := filepath.FromSlash("./internal/database/init/players.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var players []models.Player
	err = json.Unmarshal(bytes, &players)

	if err != nil {
		fmt.Println("JSON decode error!", err)
		return nil
	}
	return players
}

const insertPlayers string = `INSERT OR IGNORE INTO players (player_id, player_name, player_email, player_password, player_score, active, connected, created) VALUES (:player_id, :player_name, :player_email, :player_password, :player_score, :active, :connected, :created);`

func initPlayerTable(db *sqlx.DB) {

	players := playerTable()

	for _, p := range players {
		_, err := db.NamedExec(insertPlayers, &p)
		if err != nil {
			log.Fatalln("Error inserting player: ", err)
		}
	}

}
