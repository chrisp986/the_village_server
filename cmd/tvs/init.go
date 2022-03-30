package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
  `

const insertResources string = `INSERT OR IGNORE INTO resources (resource, quality, rate) VALUES (:resource, :quality, :rate)`

const insertVillages string = `INSERT OR IGNORE INTO villages (village_id, player_id, village_name, village_size, village_status, village_loc_y, village_loc_x) VALUES (:village_id, :player_id, :village_name, :village_size, :village_status, :village_loc_y, :village_loc_x)`

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
	}
}

func resourcesTable() []models.Resource {
	bytes, err := ioutil.ReadFile(".\\internal\\database\\init\\resources.json")

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
	bytes, err := ioutil.ReadFile(".\\internal\\database\\init\\villages.json")

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
