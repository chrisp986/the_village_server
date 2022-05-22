package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chrisp986/the_village_server/internal/database"
	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
)

type application struct {
	players interface {
		Insert(models.Player) (uint32, error)
		Get(uint32) (*models.Player, error)
	}
	villages interface {
		Insert(models.Village) (uint32, error)
	}
	villageSetup interface {
		Insert(models.VillageSetup) (uint32, error)
		InsertWithIDCheck(uint32, uint32) (uint32, error)
		GetBuildingCount(uint32) (string, error)
		UpdateBuildingString(string, models.BuildingRowAndVillage) (bool, error)
	}
	villageResources interface {
		Insert(models.VillageResource) (uint32, error)
	}
	calcResources interface {
		CalculateResources() error
	}
	buildingQueue interface {
		Insert(models.BuildingQueue) (uint32, error)
		StartConstructionNewBuilding(models.BuildingQueue) error
		UpdateBuildingQueue() ([]models.BuildingRowAndVillage, error)
		SetBuildingToDone() error
	}
}

func main() {
	log.Println("Starting the village server v0.1")

	version := flag.Int("version", 1, "genesis version")
	flag.Parse()

	err := initDB()
	if err != nil {
		log.Fatalln("Error initializing database:", err)
		return
	}

	db, err := connectDB()
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	log.Println("Connected to database")

	genesisTick(db, *version)

	initPlayerTable(db)
	// initResourceTable(db)
	initBuildingsTable(db)
	initVillageTable(db)

	resources := resourcesTable()
	buildings := buildingsTable()

	app := &application{
		players:  &database.PlayerModel{DB: db},
		villages: &database.VillageModel{DB: db},
		villageSetup: &database.VillageSetupModel{
			DB:        db,
			Buildings: buildings,
		},
		villageResources: &database.VillageResourcesModel{DB: db},
		calcResources: &database.CalcResourcesModel{
			DB:        db,
			Resources: resources,
			Buildings: buildings,
		},
		buildingQueue: &database.BuildingQueueModel{
			DB:        db,
			Resources: resources,
			Buildings: buildings,
		},
	}

	s := gocron.NewScheduler(time.UTC)
	s.StartAsync()

	if _, err := s.Every(10).Seconds().Do(app.calcResources.CalculateResources); err != nil {
		log.Println("Error in the cron job app.calcResources.CalculateResources", err)
	}

	if _, err := s.Every(1).Second().Do(func() {
		rowVillageAmount, err := app.buildingQueue.UpdateBuildingQueue()
		if err != nil {
			log.Println("Error in the cron job app.buildingQueue.UpdateBuildingQueue", err)
		}

		if len(rowVillageAmount) > 0 {
			fmt.Println("RowAndVillage: ", rowVillageAmount)
			for _, v := range rowVillageAmount {
				bString, err := app.villageSetup.GetBuildingCount(v.VillageID)
				if err != nil {
					log.Println("Error in the cron job app.villageSetup.GetBuildingCount", err)
				}
				updatedString, err := app.villageSetup.UpdateBuildingString(bString, v)
				if err != nil {
					log.Println("Error in the cron job app.villageSetup.UpdateBuildingString", err)
				}
				if updatedString {
					log.Println("Updated building string")
					err := app.buildingQueue.SetBuildingToDone()
					if err != nil {
						log.Println("Error in the cron job app.buildingQueue.SetBuildingToDone", err)
					}
				}
			}
		}

	}); err != nil {
		log.Println("Error in the cron job app.buildingQueue.UpdateBuildingQueue", err)
	}

	app.runServer()
	defer closeServer(db)
}

func closeServer(db *sqlx.DB) {

	// Close the database connection
	db.Close()
	log.Println("Closing server and delete database")
	// e := os.Remove("tv_server.db")
	// if e != nil {
	// 	log.Fatal(e)
	// }
}

func genesisTick(db *sqlx.DB, status int) {

	const genesisInsert string = `INSERT OR IGNORE INTO genesis (genesis_tick, status) VALUES (?, ?);`
	genesis := time.Now().Unix()
	log.Println("Genesis tick:", genesis)

	db.MustExec(genesisInsert, genesis, status)

}

func getGenesis(db *sqlx.DB) int32 {

	var g int32
	const getGenesis string = `SELECT * FROM genesis;`

	err := db.Get(&g, getGenesis)
	if err != nil {
		log.Fatalln("Error getting genesis:", err)
		return 0
	}
	return g
}
