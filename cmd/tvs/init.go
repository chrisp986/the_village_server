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

CREATE TABLE IF NOT EXISTS player_resources (
	village_id INTEGER PRIMARY KEY, 
	player_id INTEGER NOT NULL,
	food INTEGER NOT NULL,
	wood INTEGER NOT NULL,
	stone INTEGER NOT NULL,
	copper INTEGER NOT NULL,
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

  CREATE TABLE IF NOT EXISTS buildings (
	building_id INTEGER PRIMARY KEY,
	village_id INTEGER NOT NULL,
	player_id INTEGER NOT NULL,
	status INTEGER NOT NULL
  );

  CREATE TABLE IF NOT EXISTS resources (
	resource_id INTEGER PRIMARY KEY,
	resource TEXT NOT NULL,
	quality INTEGER NOT NULL,
	rate INTEGER NOT NULL,
	UNIQUE(resource)
  );

  CREATE TABLE IF NOT EXISTS prod_buildings_cfg (
	building_id INTEGER PRIMARY KEY,
	resource TEXT NOT NULL,
	quality INTEGER NOT NULL,
	res_rate INTEGER NOT NULL,
	res_1 INTEGER NOT NULL,
	cost_res_1 INTEGER NOT NULL,
	res_2 INTEGER NOT NULL,
	cost_res_2 INTEGER NOT NULL,
	res_3 INTEGER NOT NULL,
	cost_res_3 INTEGER NOT NULL,
	res_4 INTEGER NOT NULL,
	cost_res_4 INTEGER NOT NULL,
	res_5 INTEGER NOT NULL,
	cost_res_5 INTEGER NOT NULL,
	UNIQUE(building_id)
  );

CREATE TABLE IF NOT EXISTS village_setup (
	village_id INTEGER NOT NULL PRIMARY KEY, 
	player_id INTEGER NOT NULL,
	hunterhut_1 INTEGER NOT NULL DEFAULT 1, 
	hunterhut_2 INTEGER NOT NULL, 
	hunterhut_3 INTEGER NOT NULL, 
	hunterhut_4 INTEGER NOT NULL, 
	hunterhut_5 INTEGER NOT NULL,
	woodcutterhut_1 INTEGER NOT NULL DEFAULT 1, 
	woodcutterhut_2 INTEGER NOT NULL, 
	woodcutterhut_3 INTEGER NOT NULL, 
	woodcutterhut_4 INTEGER NOT NULL, 
	woodcutterhut_5 INTEGER NOT NULL,
	quarry_1 INTEGER NOT NULL DEFAULT 1, 
	quarry_2 INTEGER NOT NULL, 
	quarry_3 INTEGER NOT NULL, 
	quarry_4 INTEGER NOT NULL, 
	quarry_5 INTEGER NOT NULL,
	coppermine_1 INTEGER NOT NULL DEFAULT 1, 
	coppermine_2 INTEGER NOT NULL, 
	coppermine_3 INTEGER NOT NULL, 
	coppermine_4 INTEGER NOT NULL, 
	coppermine_5 INTEGER NOT NULL,
	fountain_1 INTEGER NOT NULL DEFAULT 1, 
	fountain_2 INTEGER NOT NULL, 
	fountain_3 INTEGER NOT NULL, 
	fountain_4 INTEGER NOT NULL, 
	fountain_5 INTEGER NOT NULL,
	UNIQUE(village_id)
	);
  `

const insertVillageSetup string = `INSERT OR IGNORE INTO village_setup (village_id, player_id, hunterhut_1, hunterhut_2, hunterhut_3, hunterhut_4, hunterhut_5, woodcutterhut_1, woodcutterhut_2, woodcutterhut_3, woodcutterhut_4, woodcutterhut_5,quarry_1, quarry_2, quarry_3, quarry_4, quarry_5, coppermine_1, coppermine_2, coppermine_3, coppermine_4, coppermine_5, fountain_1, fountain_2, fountain_3, fountain_4, fountain_5) VALUES (:village_id, :player_id, :hunterhut_1, :hunterhut_2, :hunterhut_3, :hunterhut_4, :hunterhut_5, :woodcutterhut_1, :woodcutterhut_2, :woodcutterhut_3, :woodcutterhut_4, :woodcutterhut_5, :quarry_1, :quarry_2, :quarry_3, :quarry_4, :quarry_5, :coppermine_1, :coppermine_2, :coppermine_3, :coppermine_4, :coppermine_5, :fountain_1, :fountain_2, :fountain_3, :fountain_4, :fountain_5)`

const insertResources string = `INSERT OR IGNORE INTO resources (resource, quality, rate) VALUES (:resource, :quality, :rate)`

const insertVillages string = `INSERT OR IGNORE INTO villages (village_id, player_id, village_name, village_size, village_status, village_loc_y, village_loc_x) VALUES (:village_id, :player_id, :village_name, :village_size, :village_status, :village_loc_y, :village_loc_x)`

const insertVillageSetupInit string = `INSERT OR IGNORE INTO village_setup (village_id, player_id, hunterhut_1, hunterhut_2, hunterhut_3, hunterhut_4, hunterhut_5, woodcutterhut_1, woodcutterhut_2, woodcutterhut_3, woodcutterhut_4, woodcutterhut_5,quarry_1, quarry_2, quarry_3, quarry_4, quarry_5, coppermine_1, coppermine_2, coppermine_3, coppermine_4, coppermine_5, fountain_1, fountain_2, fountain_3, fountain_4, fountain_5) VALUES (:village_id, :player_id, :hunterhut_1, :hunterhut_2, :hunterhut_3, :hunterhut_4, :hunterhut_5, :woodcutterhut_1, :woodcutterhut_2, :woodcutterhut_3, :woodcutterhut_4, :woodcutterhut_5, :quarry_1, :quarry_2, :quarry_3, :quarry_4, :quarry_5, :coppermine_1, :coppermine_2, :coppermine_3, :coppermine_4, :coppermine_5, :fountain_1, :fountain_2, :fountain_3, :fountain_4, :fountain_5)`

const insert string = `INSERT INTO prod_buildings_cfg (building_id,	resource, quality, res_rate, 
	res_1,cost_res_1,res_2, cost_res_2, res_3, cost_res_3, res_4, cost_res_4, res_5, cost_res_5) VALUES (:building_id, :resource, :quality, :res_rate, :res_1, :cost_res_1, :res_2, :cost_res_2, :res_3, :cost_res_3, :res_4, :cost_res_4, :res_5, :cost_res_5);`

//   CREATE TABLE IF NOT EXISTS building_info (
// 	building_id INTEGER PRIMARY KEY,
// 	resource TEXT NOT NULL,
// 	quality INTEGER NOT NULL,
// 	rate INTEGER NOT NULL,
// 	UNIQUE(resource_id)
//   );

//   CREATE TABLE IF NOT EXISTS buildings (
// 	resource_id INTEGER PRIMARY KEY,
// 	resource TEXT NOT NULL,
// 	quality INTEGER NOT NULL,
// 	rate INTEGER NOT NULL,
// 	UNIQUE(resource_id)
//   );

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

func initResourceTable(db *sqlx.DB) {

	res := resourcesTable()

	for _, r := range res {
		_, err := db.NamedExec(insertResources, &r)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func initVillageTable(db *sqlx.DB) {

	vil := villageTable()

	for _, v := range vil {
		_, err := db.NamedExec(insertVillages, &v)
		if err != nil {
			log.Fatalln(err)
		}

		vs := models.VillageSetup{
			VillageID:       v.VillageID,
			PlayerID:        v.PlayerID,
			HunterHut_1:     1,
			WoodcutterHut_1: 1,
			Quarry_1:        1,
			CopperMine_1:    1,
			Fountain_1:      1,
		}

		_, err = db.NamedExec(insertVillageSetupInit, &vs)
	}
}

func resourcesTable() []models.Resource {

	file := filepath.FromSlash("./internal/database/init/resources.json")
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Unable to load config file!", err)
		return nil
	}

	var res []models.Resource
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
		vil[i].VillageLocY = int32(randInt(0, 100))
		vil[i].VillageLocX = int32(randInt(0, 100))
	}

	return vil
}

func initPlayerResourcesTable(db *sqlx.DB) {

}
