package database

import (
	"fmt"
	"log"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type CalcResourcesModel struct {
	DB *sqlx.DB
}

func (a *CalcResourcesModel) CalculateResources() error {
	fmt.Println("Calculating resources...")

	// Ablauf:
	// 1. Get all active player ids
	// 2.

	playerIDs, err := a.GetActivePlayers()
	if err != nil {
		log.Println("Error while getting active players: ", err)
		return err
	}

	fmt.Print("Active players: ", len(playerIDs))
	if len(playerIDs) > 0 {
		fmt.Println("| Player IDs:", playerIDs)
	}

	resource, err := a.GetResourceCalcRates()
	if err != nil {
		log.Println("Error while getting resource calc rates: ", err)
		return err
	}

	err = a.GetVillageSetupFromActivePlayers(playerIDs, resource)
	if err != nil {
		log.Println("Error while getting village setup from active players: ", err)
		return err
	}

	return nil
}

func (a *CalcResourcesModel) GetActivePlayers() ([]uint32, error) {

	var playerIDs []uint32

	err := a.DB.Select(&playerIDs, "SELECT player_id FROM players WHERE active = 1;")
	if err != nil {
		return nil, err
	}

	return playerIDs, nil
}

func (a *CalcResourcesModel) GetVillageSetupFromActivePlayers(playerIDs []uint32, resource []models.Resources) error {

	for _, p := range playerIDs {

		var vs []models.VillageSetup
		err := a.DB.Select(&vs, "SELECT * FROM village_setup WHERE player_id = ?;", p)
		if err != nil {
			return err
		}

	}

	return nil
}

func (a *CalcResourcesModel) GetResourceCalcRates() ([]models.Resources, error) {

	var res []models.Resources
	err := a.DB.Select(&res, "SELECT * FROM resources;")
	if err != nil {
		return res, err
	}

	return res, nil
}
