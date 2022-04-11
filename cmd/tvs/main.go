package main

import (
	"flag"
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
	}
	villageResources interface {
		Insert(models.VillageResource) (uint32, error)
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
	defer db.Close()
	log.Println("Connected to database")

	genesisTick(db, *version)

	initResourceTable(db)
	initVillageTable(db)

	app := &application{
		players:          &database.PlayerModel{DB: db},
		villages:         &database.VillageModel{DB: db},
		villageSetup:     &database.VillageSetupModel{DB: db},
		villageResources: &database.VillageResourcesModel{DB: db}}

	s := gocron.NewScheduler(time.UTC)
	s.StartAsync()

	if _, err := s.Every(5).Seconds().Do(calcResources); err != nil {
		log.Println("Error in the cron job", err)
	}

	app.runServer()
}

func genesisTick(db *sqlx.DB, version int) {

	const genesisInsert string = `INSERT OR IGNORE INTO genesis (genesis_tick, version) VALUES (?, ?);`
	genesis := time.Now().Unix()
	log.Println("Genesis tick:", genesis)

	db.MustExec(genesisInsert, genesis, version)

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
