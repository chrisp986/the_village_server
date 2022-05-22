package models

type Player struct {
	PlayerID       uint32 `json:"player_id" db:"player_id"`
	PlayerName     string `json:"player_name" db:"player_name"`
	PlayerEmail    string `json:"player_email" db:"player_email"`
	PlayerPassword string `json:"player_password" db:"player_password"`
	PlayerScore    uint64 `json:"player_score" db:"player_score"`
	Active         uint32 `json:"active" db:"active"`
	Connected      uint32 `json:"connected" db:"connected"`
	Created        string `json:"created" db:"created"`
}

type VillageResource struct {
	VillageID uint32 `json:"village_id" db:"village_id"`
	PlayerID  uint32 `json:"player_id" db:"player_id"`
	Food      uint32 `json:"food" db:"food"`
	Wood      uint32 `json:"wood" db:"wood"`
	Stone     uint32 `json:"stone" db:"stone"`
	Metal     uint32 `json:"metal" db:"metal"`
	Water     uint32 `json:"water" db:"water"`
	Gold      uint32 `json:"gold" db:"gold"`
}

type Production struct {
	HunterHut     uint64 `json:"hunter_hut" db:"hunter_hut"`
	WoodcutterHut uint64 `json:"woodcutter_hut" db:"woodcutter_hut"`
	Quarry        uint64 `json:"quarry" db:"quarry"`
	Mine          uint64 `json:"Mine" db:"Mine"`
	Fountain      uint64 `json:"fountain" db:"fountain"`
}

type ResourceType struct {
	Food  uint64 `json:"food" db:"food"`
	Wood  uint64 `json:"wood" db:"wood"`
	Stone uint64 `json:"stone" db:"stone"`
	Metal uint64 `json:"metal" db:"metal"`
	Water uint64 `json:"water" db:"water"`
}

type Resources struct {
	ResourceID uint32 `json:"resource_id" db:"resource_id"`
	Resource   string `json:"resource" db:"resource"`
	Quality    uint32 `json:"quality" db:"quality"`
}

type ResourceRates struct {
	ResourceID uint32 `json:"resource_id" db:"resource_id"`
	Quality    uint32 `json:"quality" db:"quality"`
	Rate       uint32 `json:"rate" db:"rate"`
}

type Buildings struct {
	BuildingID     string         `json:"building_id" db:"building_id"`
	Name           string         `json:"name" db:"name"`
	Quality        uint32         `json:"quality" db:"quality"`
	ResourceID     uint32         `json:"resource_id" db:"resource_id"`
	ProductionRate uint32         `json:"production_rate" db:"production_rate"`
	BuildCost      []BuildingCost `json:"build_cost" db:"build_cost"`
	UpgradeCost    []BuildingCost `json:"upgrade_cost" db:"upgrade_cost"`
	BuildTime      uint32         `json:"build_time" db:"build_time"`
	UpgradeTime    uint32         `json:"upgrade_time" db:"upgrade_time"`
}

type BuildingCost struct {
	ResourceID uint32 `json:"resource_id" db:"resource_id"`
	Amount     uint32 `json:"amount" db:"amount"`
}

type BuildingSQL struct {
	BuildingID     string `json:"building_id" db:"building_id"`
	Name           string `json:"name" db:"name"`
	Quality        uint32 `json:"quality" db:"quality"`
	ResourceID     uint32 `json:"resource_id" db:"resource_id"`
	ProductionRate uint32 `json:"production_rate" db:"production_rate"`
	BuildCost      string `json:"build_cost" db:"build_cost"`
	UpgradeCost    string `json:"upgrade_cost" db:"upgrade_cost"`
	BuildTime      uint32 `json:"build_time" db:"build_time"`
	UpgradeTime    uint32 `json:"upgrade_time" db:"upgrade_time"`
}

type BuildingCount struct {
	BuildingID string `json:"building_id" db:"building_id"`
	Count      uint32 `json:"count" db:"count"`
}

type BuildingRowAndVillage struct {
	RowID      uint32 `json:"rowid" db:"rowid"`
	BuildingID string `json:"building_id" db:"building_id"`
	VillageID  uint32 `json:"village_id" db:"village_id"`
	Amount     uint32 `json:"amount" db:"amount"`
}

type BuildingConfig struct {
	BuildingID    uint32 `json:"building_id" db:"building_id"`
	Resource      string `json:"resource" db:"resource"`
	Quality       uint32 `json:"quality" db:"quality"`
	ResRate       uint32 `json:"res_rate" db:"res_rate"`
	Resource1     uint32 `json:"resource1" db:"resource1"`
	CostResource1 uint32 `json:"cost_res_1" db:"cost_res_1"`
	Resource2     uint32 `json:"res_2" db:"res_2"`
	CostResource2 uint32 `json:"cost_res_2" db:"cost_res_2"`
	Resource3     uint32 `json:"res_3" db:"res_3"`
	CostResource3 uint32 `json:"cost_res_3" db:"cost_res_3"`
	Resource4     uint32 `json:"res_4" db:"res_4"`
	CostResource4 uint32 `json:"cost_res_4" db:"cost_res_4"`
	Resource5     uint32 `json:"res_5" db:"res_5"`
	CostResource5 uint32 `json:"cost_res_5" db:"cost_res_5"`
}

// village_id INTEGER PRIMARY KEY,
// 	player_id INTEGER NOT NULL,
// 	village_name TEXT NOT NULL,
// 	size INTEGER NOT NULL,
// 	status INTEGER NOT NULL,
// 	location_y INTEGER NOT NULL,
// 	location_x INTEGER NOT NULL,
// 	UNIQUE(village_id)
type Village struct {
	VillageID     uint32 `json:"village_id" db:"village_id"`
	PlayerID      uint32 `json:"player_id" db:"player_id"`
	VillageName   string `json:"village_name" db:"village_name"`
	VillageSize   uint32 `json:"village_size" db:"village_size"`
	VillageStatus uint32 `json:"village_status" db:"village_status"`
	VillageLocY   uint32 `json:"village_loc_y" db:"village_loc_y"`
	VillageLocX   uint32 `json:"village_loc_x" db:"village_loc_x"`
}

// HunterHut     int64 `json:"hunter_hut" db:"hunter_hut"`
// 	WoodcutterHut int64 `json:"woodcutter_hut" db:"woodcutter_hut"`
// 	Quarry        int64 `json:"quarry" db:"quarry"`
// 	CopperMine    int64 `json:"copper_mine" db:"copper_mine"`
// 	Fountain      int64 `json:"fountain" db:"fountain"`

type VillageSetup struct {
	VillageID  uint32 `json:"village_id" db:"village_id"`
	PlayerID   uint32 `json:"player_id" db:"player_id"`
	Buildings  string `json:"buildings" db:"buildings"`
	Status     uint32 `json:"status" db:"status"`
	LastUpdate string `json:"last_update" db:"last_update"`
}

type BuildingQueue struct {
	VillageID  uint32 `json:"village_id" db:"village_id"`
	PlayerID   uint32 `json:"player_id" db:"player_id"`
	BuildingID string `json:"building_id" db:"building_id"`
	Amount     uint32 `json:"amount" db:"amount"`
	Status     uint8  `json:"status" db:"status"`
	StartTime  uint32 `json:"start_time" db:"start_time"`
	FinishTime uint32 `json:"finish_time" db:"finish_time"`
}

// CREATE TABLE IF NOT EXISTS prod_buildings_cfg (
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
