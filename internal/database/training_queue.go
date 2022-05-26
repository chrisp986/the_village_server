package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/jmoiron/sqlx"
)

type TrainingQueueModel struct {
	DB        *sqlx.DB
	Resources []models.Resources
	Workers   []models.Workers
}

const status10 uint8 = 10
const status20 uint8 = 20

// CREATE TABLE IF NOT EXISTS workers_queue (
// 	worker_id TEXT NOT NULL,
// 	village_id INTEGER NOT NULL,
// 	player_id INTEGER NOT NULL,
// 	amount INTEGER NOT NULL,
// 	status INTEGER NOT NULL,
// 	start_time TEXT NOT NULL,
// 	finish_time TEXT NOT NULL
// 	);

func (m *TrainingQueueModel) Insert(newworkers models.TrainingQueue) (uint32, error) {

	stmt := `INSERT INTO training_queue (worker_id, village_id, player_id, amount, status, start_time, finish_time)
	VALUES (:worker_id, :village_id, :player_id, :amount, :status, :start_time, :finish_time);`

	newworkers.StartTime = uint32(time.Now().Unix())

	result, err := m.DB.NamedExec(stmt, &newworkers)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

// Main function
func (m *TrainingQueueModel) StartTrainingNewWorker(trainingQueue models.TrainingQueue) error {

	workers := m.getworkersData(trainingQueue.WorkerID)

	resCheck, err := m.checkIfSufficientResources(trainingQueue, workers)
	if err != nil {
		log.Println("Error checking if sufficient resources: ", err)
		return err
	}

	if resCheck {
		m.insertTotrainingQueue(trainingQueue, workers)
	}
	return err
}

//
//TODO Create new function to keep track of the progress of the workers queue
//
func (m *TrainingQueueModel) UpdateTrainingQueue() ([]models.BuildingRowAndVillage, error) {

	// set status of the workers in the queue to start = 10 or 20

	err := m.setworkersToStart()
	if err != nil {
		log.Println("Error setting workers to start 'UpdatetrainingQueue': ", err)
		return nil, err
	}

	// check if the finish_time is reached
	// if reached, get the rowid and village_id to add the workers to the village (update the village_setup table)

	rowVillageAmount, err := m.getRowAndVillageIDs()
	if err != nil {
		log.Println("Error getting row and village IDs 'UpdatetrainingQueue': ", err)
		return nil, err
	}

	return rowVillageAmount, err
}

func (m *TrainingQueueModel) getRowAndVillageIDs() ([]models.BuildingRowAndVillage, error) {

	var data []models.BuildingRowAndVillage

	err := m.DB.Select(&data, "SELECT rowid, worker_id, village_id, amount FROM training_queue WHERE (status = ? OR status = ?) AND strftime('%s', 'now') >= finish_time;", status10, status20)

	if err == sql.ErrNoRows {
		err = nil
	}

	return data, err
}

func (m *TrainingQueueModel) SetTrainingToDone() error {

	stmt := "UPDATE training_queue SET status = status * 10 WHERE (status = ? OR status = ?) AND strftime('%s', 'now') >= finish_time;"

	result, err := m.DB.Exec(stmt, status10, status20)
	if err != nil {
		log.Println("Error setting workers to done 'setworkersToDone': ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected 'setworkersToDone': ", err)
		return err
	}

	if rowsAffected > 0 {
		log.Println("workers process completed!")
		return err
	}
	log.Println("Error! No Rows updated!")
	return err
}

func (m *TrainingQueueModel) setworkersToStart() error {

	const stmt string = "UPDATE training_queue SET status = status * 10 WHERE (status = 1 OR status = 2) AND start_time <= strftime('%s', 'now');"

	_, err := m.DB.Exec(stmt)
	if err != nil {
		log.Println("Error setting workers to start 'setworkersToStart': ", err)
		return err
	}

	return nil
}

func (m *TrainingQueueModel) getworkersData(workersID string) models.Workers {

	// var workers models.workersSQL

	for _, b := range m.Workers {
		if b.WorkerID == workersID {
			return b
		}
	}

	// stmt := fmt.Sprintf("SELECT * FROM workerss WHERE worker_id='%s';", workersID)

	// err := m.DB.Get(&workers, stmt)
	// if err != nil {
	// 	return workers, err
	// }

	return models.Workers{}
}

func (m *TrainingQueueModel) checkIfSufficientResources(trainingQueue models.TrainingQueue, workers models.Workers) (bool, error) {

	//check if the village has sufficient resources

	// bcs := splitCostString(workers.BuildCost)

	for _, bc := range workers.TrainingCost {

		// get the amount of the resource
		// get the resourceName
		resName, err := m.getResourceName(bc.ResourceID)
		if err != nil {
			log.Println("Error getting resource name: ", err)
			return false, err
		}

		resCheck, err := m.checkResourceFromVillage(resName, bc.Amount, trainingQueue.VillageID)
		if err != nil {
			log.Println("Error checking resource from village: ", err)
			return false, err
		}

		if resCheck {
			resUpdated, err := m.updateVillageResources(resName, bc.Amount, trainingQueue.VillageID)
			if err != nil {
				return false, err
			}
			if resUpdated {
				log.Printf("Updated village resources! ResourceName: %s, Cost: %d, VillageID: %d", resName, bc.Amount, trainingQueue.VillageID)

			}
		}
	}
	return true, nil
}

func (m *TrainingQueueModel) insertTotrainingQueue(trainingQueue models.TrainingQueue, workers models.Workers) (bool, error) {

	//TODO insert to workers queue
	//workersID, villageID, playerID, amount, status, startTime, finishTime

	const inserttrainingQueue string = `INSERT INTO training_queue (worker_id, village_id, player_id, amount, status, start_time, finish_time) VALUES (:worker_id, :village_id, :player_id, :amount, :status, :start_time, :finish_time)`

	trainingQueue.StartTime = uint32(time.Now().Unix())
	trainingQueue.FinishTime = trainingQueue.StartTime + trainingQueue.Amount*workers.TrainingTime

	tx := m.DB.MustBegin()
	_, err := tx.NamedExec(inserttrainingQueue, &trainingQueue)
	switch err {
	case nil:
		tx.Commit()
		log.Printf("Inserted workers queue! workersID: %s, VillageID: %d, PlayerID: %d, Amount: %d, Status: %d, StartTime: %d, FinishTime: %d", trainingQueue.WorkerID, trainingQueue.VillageID, trainingQueue.PlayerID, trainingQueue.Amount, trainingQueue.Status, trainingQueue.StartTime, trainingQueue.FinishTime)
		return true, err
	default:
		log.Printf("Error inserting to workers queue: %s, workersID: %s, VillageID: %d, PlayerID: %d, Amount: %d, Status: %d, StartTime: %d, FinishTime: %d", err, trainingQueue.WorkerID, trainingQueue.VillageID, trainingQueue.PlayerID, trainingQueue.Amount, trainingQueue.Status, trainingQueue.StartTime, trainingQueue.FinishTime)
		tx.Rollback()
		return false, err
	}
}

func (m *TrainingQueueModel) updateVillageResources(resName string, resCost uint32, villageID uint32) (bool, error) {

	stmt := fmt.Sprintf("UPDATE village_resources SET %s = %s - %d WHERE village_id=%d;", resName, resName, resCost, villageID)

	tx := m.DB.MustBegin()
	_, err := tx.Exec(stmt)
	switch err {
	case nil:
		tx.Commit()
		return true, err
	default:
		log.Printf("Error! Rollback updating village resources: %s - resName: %s, resCost: %d, villageID: %d", err, resName, resCost, villageID)
		tx.Rollback()
		return false, err
	}
}

func (m *TrainingQueueModel) checkResourceFromVillage(resourceName string, resCost uint32, villageID uint32) (bool, error) {

	// returns 1 if the village has enough resources and 0 if no sufficient resources

	var resCheck uint8
	stmt := fmt.Sprintf("SELECT CASE WHEN %s >= %d THEN '1' ELSE '0' END FROM village_resources WHERE village_id=%d;", resourceName, resCost, villageID)

	err := m.DB.Get(&resCheck, stmt)
	if err != nil {
		log.Println("Error getting resource amount: ", err)
		return false, err
	}

	if resCheck == 1 {
		return true, nil
	}

	return false, nil
}

func (m *TrainingQueueModel) getResourceName(resourceID uint32) (string, error) {

	for _, r := range m.Resources {
		if r.ResourceID == resourceID {
			return r.Resource, nil
		}
	}

	return "", nil
}

func splitCostString(s string) []models.TrainingCost {

	var bcs []models.TrainingCost

	s1 := strings.Split(s, ",")

	for _, v := range s1 {
		if v != "" {

			res := matchResource(v)
			if res == 4294967295 {
				log.Println("Error: Resource not in range")
			}
			amount := matchAmount(v)
			if amount == 4294967295 {
				log.Println("Error: Amount not in range")
			}

			bcs = append(bcs, models.TrainingCost{ResourceID: res, Amount: amount})
		}
	}
	return bcs
}

func matchResource(s string) uint32 {

	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
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

func matchAmount(s string) uint32 {

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
