package database

import (
	"fmt"
	"log"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type CalcResourcesModel struct {
	DB        *sqlx.DB
	Resources []models.Resources
	Workers   []models.Workers
}

func (a *CalcResourcesModel) CalculateResources() error {

	// Ablauf:
	// 1. Get all active player ids
	// 2. Get the buildingsString from each village_setup for player_id (loop)
	// What is needed to calculate the new resource:
	// The old resource amount -> village_resources
	// The amount of buildings from that specific building type -> village_setup.buildings, village_setup.village_id
	// The calc_rate from the resources table -> resources.calc_rate

	// village_id, player_id, building_id, building_count, calc_rate, old_resource_amount

	playerIDs := a.getActivePlayers()

	if len(playerIDs) > 0 {
		log.Println("Active players:", len(playerIDs))

	}

	// for _, pID := range playerIDs {
	// 	villages := a.getVillagesFromActivePlayer(pID)

	// for _, v := range villages {

	// buildingString := a.getBuildingString(v.VillageID, v.PlayerID)
	// bcs := splitString(buildingString)

	// for _, b := range bcs {
	// 	if b.Count > 0 {

	// 		buildings := a.getResourceRate(b.BuildingID)
	// 		// log.Println("Rate: ", buildings.ProductionRate)
	// 		// resources := a.getVillageResources(v.VillageID, v.PlayerID)
	// 		// log.Println("Resource Food: ", resources.Food)
	// 		oldResCount, att := a.getOldResourceCount(buildings.ResourceID, v.VillageID, v.PlayerID)

	// 		newResCount := oldResCount + (buildings.ProductionRate * b.Count)

	// 		log.Println("VillageID: ", v.VillageID, " PlayerID: ", v.PlayerID, " BuildingID: ", b.BuildingID, " Amount of buildings: ", b.Count, " Product: ", att, " Rate: ", buildings.ProductionRate, " Old Resource Count: ", oldResCount, " New Resource Count: ", newResCount)

	// 		// Update village_resources
	// 		err := a.updateVillageResources(v.VillageID, v.PlayerID, att, newResCount)
	// 		if err != nil {
	// 			log.Println("Error: ", err)
	// 		}
	// 	}
	return nil
}

// resource, err := a.GetResourceCalcRates()
// if err != nil {
// 	log.Println("Error while getting resource calc rates: ", err)
// 	return err
// }

// err = a.GetVillageSetupFromActivePlayers(playerIDs, resource)
// if err != nil {
// 	log.Println("Error while getting village setup from active players: ", err)
// 	return err
// }

// 	return nil
// }

func (a *CalcResourcesModel) getActivePlayers() []uint32 {

	var playerIDs []uint32

	err := a.DB.Select(&playerIDs, "SELECT player_id FROM players WHERE active = 1;")
	if err != nil {
		log.Println("Error while getting active players: ", err)
		return nil
	}

	return playerIDs
}

func (a *CalcResourcesModel) getVillagesFromActivePlayer(playerIDs uint32) []models.Village {

	var v []models.Village
	err := a.DB.Select(&v, "SELECT * FROM villages WHERE player_id = ?;", playerIDs)
	if err != nil {
		log.Println("Error while getting villages from active player: ", err, " PlayerID: ", playerIDs)
		return nil
	}

	return v
}

func (a *CalcResourcesModel) getBuildingString(villageID uint32, playerID uint32) string {

	var buildings string

	err := a.DB.Get(&buildings, "SELECT buildings FROM village_setup WHERE village_id=? AND player_id=?;", villageID, playerID)
	if err != nil {
		log.Println("Error while getting building string: ", err, " VillageID: ", villageID, " PlayerID: ", playerID)
		return ""
	}

	return buildings
}

func (a *CalcResourcesModel) getResourceRate(workerID string) models.Workers {

	for _, w := range a.Workers {
		if w.WorkerID == workerID {
			return w
		}

	}
	log.Println("Error while getting resource rate: ", workerID)
	return models.Workers{}
}

func (a *CalcResourcesModel) getVillageResources(villageID uint32, playerID uint32) models.VillageResource {

	var villageRes models.VillageResource

	err := a.DB.Get(&villageRes, "SELECT * FROM village_resources WHERE village_id=? AND player_id=?;", villageID, playerID)
	if err != nil {
		log.Println("Error while getting village resources: ", err, " VillageID: ", villageID, " PlayerID: ", playerID)
		return villageRes
	}

	return villageRes
}

func (a *CalcResourcesModel) getOldResourceCount(resourceID uint32, villageID uint32, playerID uint32) (uint32, string) {

	//Identify which resource needs to be calculated -> switch case

	var att string
	switch resourceID {
	case 0:
		// Food
		att = "food"
	case 1:
		// Wood
		att = "wood"
	case 2:
		// Stone
		att = "stone"
	case 3:
		// Metal
		att = "Metal"
	case 4:
		// Water
		att = "water"
	case 5:
		// Gold
		att = "gold"
	}

	stmt := fmt.Sprintf("SELECT %s FROM village_resources WHERE village_id = ? AND player_id = ?;", att)

	var oldResCount uint32

	err := a.DB.Get(&oldResCount, stmt, villageID, playerID)
	if err != nil {
		log.Println("Error while getting old resource count: ", err, " VillageID: ", villageID, " PlayerID: ", playerID, " Attribute: ", att)
		return 0, ""
	}

	return oldResCount, att
}

func (a *CalcResourcesModel) updateVillageResources(villageID uint32, playerID uint32, att string, newResCount uint32) error {

	stmt := fmt.Sprintf("UPDATE village_resources SET %s = ? WHERE village_id = ? AND player_id = ?;", att)

	tx := a.DB.MustBegin()
	_, err := tx.Exec(stmt, newResCount, villageID, playerID)
	switch err {
	case nil:
		tx.Commit()
	default:
		log.Println("Error while updating village resources: ", err, " VillageID: ", villageID, " PlayerID: ", playerID, " Attribute: ", att)
		tx.Rollback()
	}

	return err
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

// func (a *CalcResourcesModel) GetResourceCalcRates() ([]models.Resources, error) {

// 	var res []models.Resources
// 	err := a.DB.Select(&res, "SELECT * FROM resources;")
// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }

// func splitString(s string) []models.BuildingCount {

// 	var bcs []models.BuildingCount

// 	s1 := strings.Split(s, ",")

// 	for _, v := range s1[:len(s1)-1] {
// 		s2 := strings.Split(v, "=")
// 		b := s2[0]
// 		c := s2[1]

// 		c64, err := strconv.ParseUint(c, 10, 32)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		bcs = append(bcs, models.BuildingCount{
// 			WorkerID: b,
// 			Count:    uint32(c64),
// 		})

// 	}

// 	return bcs
// }

// func SplitBuildingsString(s string) []models.BuildingCount {

// 	var bcs []models.BuildingCount

// 	s1 := strings.Split(s, ",")

// 	for _, v := range s1 {
// 		if v != "" {

// 			b := matchB(v)
// 			if b == "" {
// 				log.Println("Error: Building ID not in range")
// 			}
// 			c := matchC(v)
// 			if c == 4294967295 {
// 				log.Println("Error: Count not in range")
// 			}

// 			bcs = append(bcs, models.BuildingCount{
// 				BuildingID: b,
// 				Count:      c,
// 			})
// 		}
// 	}
// 	return bcs
// }

// func matchB(s string) string {

// 	i := strings.Index(s, "(")
// 	if i >= 0 {
// 		j := strings.Index(s, ")")
// 		if j >= 0 {
// 			b := s[i+1 : j]
// 			return b
// 		}
// 	}
// 	return ""
// }

// func matchC(s string) uint32 {

// 	i := strings.Index(s, "[")
// 	if i >= 0 {
// 		j := strings.Index(s, "]")
// 		if j >= 0 {
// 			c := s[i+1 : j]
// 			c64, err := strconv.ParseUint(c, 10, 32)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			return uint32(c64)
// 		}
// 	}
// 	return 4294967295
// }
