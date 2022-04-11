package models

type Player struct {
	PlayerID       uint32 `json:"player_id" db:"player_id"`
	PlayerName     string `json:"player_name" db:"player_name"`
	PlayerEmail    string `json:"player_email" db:"player_email"`
	PlayerPassword string `json:"player_password" db:"player_password"`
	PlayerScore    uint64 `json:"player_score" db:"player_score"`
	Active         bool   `json:"active" db:"active"`
	Connected      bool   `json:"connected" db:"connected"`
	Created        string `json:"created" db:"created"`
}

type VillageResource struct {
	VillageID uint32 `json:"village_id" db:"village_id"`
	PlayerID  uint32 `json:"player_id" db:"player_id"`
	Food      uint32 `json:"food" db:"food"`
	Wood      uint32 `json:"wood" db:"wood"`
	Stone     uint32 `json:"stone" db:"stone"`
	Copper    uint32 `json:"copper" db:"copper"`
	Water     uint32 `json:"water" db:"water"`
	Gold      uint32 `json:"gold" db:"gold"`
}

type Production struct {
	HunterHut     uint64 `json:"hunter_hut" db:"hunter_hut"`
	WoodcutterHut uint64 `json:"woodcutter_hut" db:"woodcutter_hut"`
	Quarry        uint64 `json:"quarry" db:"quarry"`
	CopperMine    uint64 `json:"copper_mine" db:"copper_mine"`
	Fountain      uint64 `json:"fountain" db:"fountain"`
}

type ResourceType struct {
	Food   uint64 `json:"food" db:"food"`
	Wood   uint64 `json:"wood" db:"wood"`
	Stone  uint64 `json:"stone" db:"stone"`
	Copper uint64 `json:"copper" db:"copper"`
	Water  uint64 `json:"water" db:"water"`
}

type Resource struct {
	ResourceID uint32 `json:"resource_id" db:"resource_id"`
	Resource   string `json:"resource" db:"resource"`
	Quality    uint32 `json:"quality" db:"quality"`
	Rate       uint32 `json:"rate" db:"rate"`
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
	VillageID       uint32 `json:"village_id" db:"village_id"`
	PlayerID        uint32 `json:"player_id" db:"player_id"`
	HunterHut_1     uint32 `json:"hunterhut_1" db:"hunterhut_1"`
	HunterHut_2     uint32 `json:"hunterhut_2" db:"hunterhut_2"`
	HunterHut_3     uint32 `json:"hunterhut_3" db:"hunterhut_3"`
	HunterHut_4     uint32 `json:"hunterhut_4" db:"hunterhut_4"`
	HunterHut_5     uint32 `json:"hunterhut_5" db:"hunterhut_5"`
	WoodcutterHut_1 uint32 `json:"woodcutterhut_1" db:"woodcutterhut_1"`
	WoodcutterHut_2 uint32 `json:"woodcutterhut_2" db:"woodcutterhut_2"`
	WoodcutterHut_3 uint32 `json:"woodcutterhut_3" db:"woodcutterhut_3"`
	WoodcutterHut_4 uint32 `json:"woodcutterhut_4" db:"woodcutterhut_4"`
	WoodcutterHut_5 uint32 `json:"woodcutterhut_5" db:"woodcutterhut_5"`
	Quarry_1        uint32 `json:"quarry_1" db:"quarry_1"`
	Quarry_2        uint32 `json:"quarry_2" db:"quarry_2"`
	Quarry_3        uint32 `json:"quarry_3" db:"quarry_3"`
	Quarry_4        uint32 `json:"quarry_4" db:"quarry_4"`
	Quarry_5        uint32 `json:"quarry_5" db:"quarry_5"`
	CopperMine_1    uint32 `json:"coppermine_1" db:"coppermine_1"`
	CopperMine_2    uint32 `json:"coppermine_2" db:"coppermine_2"`
	CopperMine_3    uint32 `json:"coppermine_3" db:"coppermine_3"`
	CopperMine_4    uint32 `json:"coppermine_4" db:"coppermine_4"`
	CopperMine_5    uint32 `json:"coppermine_5" db:"coppermine_5"`
	Fountain_1      uint32 `json:"fountain_1" db:"fountain_1"`
	Fountain_2      uint32 `json:"fountain_2" db:"fountain_2"`
	Fountain_3      uint32 `json:"fountain_3" db:"fountain_3"`
	Fountain_4      uint32 `json:"fountain_4" db:"fountain_4"`
	Fountain_5      uint32 `json:"fountain_5" db:"fountain_5"`
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
