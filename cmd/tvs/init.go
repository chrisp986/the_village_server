package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"time"

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

CREATE TABLE IF NOT EXISTS training_queue (
	worker_id TEXT NOT NULL,
	village_id INTEGER NOT NULL,
	player_id INTEGER NOT NULL,
	amount INTEGER NOT NULL,
	status INTEGER NOT NULL,
	start_time TEXT NOT NULL,
	finish_time TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS village_setup (village_id INTEGER PRIMARY KEY, player_id INTEGER NOT NULL,'h1' INTEGER NOT NULL,'h2' INTEGER NOT NULL,'h3' INTEGER NOT NULL,'h4' INTEGER NOT NULL,'h5' INTEGER NOT NULL,'h6' INTEGER NOT NULL,'l1' INTEGER NOT NULL,'l2' INTEGER NOT NULL,'l3' INTEGER NOT NULL,'l4' INTEGER NOT NULL,'l5' INTEGER NOT NULL,'l6' INTEGER NOT NULL,'m1' INTEGER NOT NULL,'m2' INTEGER NOT NULL,'m3' INTEGER NOT NULL,'m4' INTEGER NOT NULL,'m5' INTEGER NOT NULL,'m6' INTEGER NOT NULL,'b1' INTEGER NOT NULL,'b2' INTEGER NOT NULL,'b3' INTEGER NOT NULL,'b4' INTEGER NOT NULL,'b5' INTEGER NOT NULL,'b6' INTEGER NOT NULL, status INTEGER NOT NULL, last_update TEXT NOT NULL);

  `

//   CREATE TABLE IF NOT EXISTS village_setup (
// 	village_id INTEGER NOT NULL PRIMARY KEY,
// 	player_id INTEGER NOT NULL,
// 	Workers TEXT NOT NULL,
// 	status INTEGER NOT NULL,
// 	last_update TEXT NOT NULL,
// 	UNIQUE(village_id)
// 	);

//   CREATE TABLE IF NOT EXISTS resources (
// 	resource_id INTEGER PRIMARY KEY,
// 	resource TEXT NOT NULL,
// 	quality INTEGER NOT NULL
//   );

//   CREATE TABLE IF NOT EXISTS Workers (
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

//   CREATE TABLE IF NOT EXISTS prod_Workers_cfg (
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

const insertVillageSetup string = `INSERT OR IGNORE INTO village_setup (village_id, player_id, worker_id, status, last_update) VALUES (:village_id, :player_id, :Workers, :status, :last_update)`

const insertResourceRates string = `INSERT OR IGNORE INTO resource_rates (resource, quality, rate) VALUES (:resource, :quality, :rate)`

const insertVillages string = `INSERT OR IGNORE INTO villages (village_id, player_id, village_name, village_size, village_status, village_loc_y, village_loc_x) VALUES (:village_id, :player_id, :village_name, :village_size, :village_status, :village_loc_y, :village_loc_x)`

const insertVillageResourcesInit string = `INSERT OR IGNORE INTO village_resources (village_id, player_id, food, wood, stone, metal, water, gold) VALUES (:village_id, :player_id, :food, :wood, :stone, :metal, :water, :gold);`

// const insert string = `INSERT INTO prod_Workers_cfg (building_id,	resource, quality, res_rate,
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

	return db, err
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

// func initVillageSetupTable(db *sqlx.DB, village models.Village, Workers []models.Wo) {
// 	//TODO

// 	insertVillageSetupInit := fmt.Sprintf("INSERT OR IGNORE INTO village_setup (village_id, player_id, Workers, status, last_update) VALUES (:village_id, :player_id, :Workers, :status, :last_update);")

// 	vs := models.VillageSetup{
// 		VillageID:  village.VillageID,
// 		PlayerID:   village.PlayerID,
// 		Workers:  database.InitWorkersString(buildingIDs),
// 		Status:     0,
// 		LastUpdate: time.Now().Local().Format("2006-01-02 15:04:05"),
// 	}

// 	_, err := db.NamedExec(insertVillageSetupInit, &vs)
// 	if err != nil {
// 		log.Fatalf("Error inserting village setup: %v, VillageID: %d, PlayerID: %d", err, village.VillageID, village.PlayerID)
// 	}
// }

func createVillageSetupTable(db *sqlx.DB, workers []models.Workers) error {

	var new string

	for _, w := range workers {
		new += fmt.Sprintf("'%s' INTEGER NOT NULL,", w.WorkerID)
	}

	villageSetupTable := "CREATE TABLE IF NOT EXISTS village_setup (village_id INTEGER PRIMARY KEY, player_id INTEGER NOT NULL," + new + " status INTEGER NOT NULL, last_update TEXT NOT NULL);"

	fmt.Println(villageSetupTable)

	if _, err := db.Exec(villageSetupTable); err != nil {
		log.Println("Error creating village setup table: ", err)
		return err
	}

	return nil
}

func initVillageResourcesTable(db *sqlx.DB, village models.Village) {

	vr := models.VillageResource{
		VillageID: village.VillageID,
		PlayerID:  village.PlayerID,
		Food:      100,
		Wood:      100,
		Stone:     100,
		Metal:     100,
		Water:     100,
		Gold:      20,
	}

	_, err := db.NamedExec(insertVillageResourcesInit, &vr)
	if err != nil {
		log.Fatalf("Error inserting village resources: %v, VillageID: %d, PlayerID: %d", err, village.VillageID, village.PlayerID)
	}
}

func initVillageTable(db *sqlx.DB) {

	vil := villageTable()

	// bID, err := getWorkersID(db)
	var wID []string
	workers := workersTable()

	for _, w := range workers {
		wID = append(wID, w.WorkerID)
	}

	for _, v := range vil {

		_, err := db.NamedExec(insertVillages, &v)
		if err != nil {
			log.Fatalln("Error inserting village: ", err)
		}

		// initVillageSetupTable(db, v, Workers)
		// initVillageSetupTable(db, v, bID)

		initVillageResourcesTable(db, v)

	}
}

// const insertWorkers string = `INSERT OR IGNORE INTO Workers (building_id, name, quality, resource_id, production_rate, build_cost, upgrade_cost, build_time, upgrade_time) VALUES (:building_id, :name, :quality, :resource_id, :production_rate, :build_cost, :upgrade_cost, :build_time, :upgrade_time);`

// func initWorkersTable(db *sqlx.DB) {

// 	WorkersJSON := WoTable()

// 	// log.Println(WorkersJSON)

// 	var Workers []models.WoQL

// 	for _, b := range WorkersJSON {
// 		Workers = append(Workers, models.WorkersQL{
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

// 	for _, b := range Workers {

// 		_, err := db.NamedExec(insertWorkers, &b)
// 		if err != nil {
// 			log.Fatalln("Error inserting building: ", err)
// 		}

// 	}
// }

func costToString(tc []models.TrainingCost) string {

	var cost string
	for _, v := range tc {
		cost += fmt.Sprintf("(%d)[%d],", v.ResourceID, v.Amount)
	}

	return cost
}

func workersTable() []models.Workers {

	file := filepath.FromSlash("./internal/database/init/workers.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var workers []models.Workers
	err = json.Unmarshal(bytes, &workers)

	if err != nil {
		fmt.Println("JSON decode error in 'WorkersTable'", err)
		return nil
	}
	return workers
}

// func getWorkersID(db *sqlx.DB) ([]string, error) {

// 	var bID []string
// 	err := db.Select(&bID, "SELECT building_id FROM Workers ORDER BY rowid;")

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
