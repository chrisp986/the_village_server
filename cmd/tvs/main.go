package main

import (
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
		Delete(uint32) error
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
	trainingQueue interface {
		Insert(models.TrainingQueue) (uint32, error)
		StartTrainingNewWorker(models.TrainingQueue) error
		UpdateTrainingQueue() ([]models.BuildingRowAndVillage, error)
		SetTrainingToDone() error
	}
}

func main() {
	log.Println("Starting the village server v0.1")

	// version := flag.Int("version", 1, "genesis version")
	// flag.Parse()

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

	// genesisTick(db, *version)

	// initResourceTable(db)
	// initBuildingsTable(db)
	initPlayerTable(db)
	initVillageTable(db)

	resources := resourcesTable()
	workers := workersTable()

	app := &application{
		players:  &database.PlayerModel{DB: db},
		villages: &database.VillageModel{DB: db},
		villageSetup: &database.VillageSetupModel{
			DB:      db,
			Workers: workers,
		},
		villageResources: &database.VillageResourcesModel{DB: db},
		calcResources: &database.CalcResourcesModel{
			DB:        db,
			Resources: resources,
			Workers:   workers,
		},
		trainingQueue: &database.TrainingQueueModel{
			DB:        db,
			Resources: resources,
			Workers:   workers,
		},
	}

	s := gocron.NewScheduler(time.UTC)
	s.StartAsync()

	if _, err := s.Every(10).Seconds().Do(app.calcResources.CalculateResources); err != nil {
		log.Println("Error in the cron job app.calcResources.CalculateResources", err)
	}

	if _, err := s.Every(1).Second().Do(func() {
		rowVillageAmount, err := app.trainingQueue.UpdateTrainingQueue()
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
					err := app.trainingQueue.SetTrainingToDone()
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
	log.Println("Closing server")
	// e := os.Remove("tv_server.db")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// log.Println("Database deleted")
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
