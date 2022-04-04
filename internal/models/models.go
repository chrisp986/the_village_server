package models

type Player struct {
	PlayerID       int32  `json:"player_id" db:"player_id"`
	PlayerName     string `json:"player_name" db:"player_name"`
	PlayerEmail    string `json:"player_email" db:"player_email"`
	PlayerPassword string `json:"player_password" db:"player_password"`
	PlayerScore    int64  `json:"player_score" db:"player_score"`
	Active         bool   `json:"active" db:"active"`
	Connected      bool   `json:"connected" db:"connected"`
	Created        string `json:"created" db:"created"`
}

type PlayerResource struct {
	VillageID int32 `json:"village_id" db:"village_id"`
	PlayerID  int32 `json:"player_id" db:"player_id"`
	Food      int32 `json:"food" db:"food"`
	Wood      int32 `json:"wood" db:"wood"`
	Stone     int32 `json:"stone" db:"stone"`
	Copper    int32 `json:"copper" db:"copper"`
	Water     int32 `json:"water" db:"water"`
	Gold      int32 `json:"gold" db:"gold"`
}

type Production struct {
	HunterHut     int64 `json:"hunter_hut" db:"hunter_hut"`
	WoodcutterHut int64 `json:"woodcutter_hut" db:"woodcutter_hut"`
	Quarry        int64 `json:"quarry" db:"quarry"`
	CopperMine    int64 `json:"copper_mine" db:"copper_mine"`
	Fountain      int64 `json:"fountain" db:"fountain"`
}

type ResourceType struct {
	Food   int64 `json:"food" db:"food"`
	Wood   int64 `json:"wood" db:"wood"`
	Stone  int64 `json:"stone" db:"stone"`
	Copper int64 `json:"copper" db:"copper"`
	Water  int64 `json:"water" db:"water"`
}

type Resource struct {
	ResourceID int32  `json:"resource_id" db:"resource_id"`
	Resource   string `json:"resource" db:"resource"`
	Quality    int32  `json:"quality" db:"quality"`
	Rate       int32  `json:"rate" db:"rate"`
}

type BuildingConfig struct {
	BuildingID    int32  `json:"building_id" db:"building_id"`
	Resource      string `json:"resource" db:"resource"`
	Quality       int32  `json:"quality" db:"quality"`
	ResRate       int32  `json:"res_rate" db:"res_rate"`
	Resource1     int32  `json:"resource1" db:"resource1"`
	CostResource1 int32  `json:"cost_res_1" db:"cost_res_1"`
	Resource2     int32  `json:"res_2" db:"res_2"`
	CostResource2 int32  `json:"cost_res_2" db:"cost_res_2"`
	Resource3     int32  `json:"res_3" db:"res_3"`
	CostResource3 int32  `json:"cost_res_3" db:"cost_res_3"`
	Resource4     int32  `json:"res_4" db:"res_4"`
	CostResource4 int32  `json:"cost_res_4" db:"cost_res_4"`
	Resource5     int32  `json:"res_5" db:"res_5"`
	CostResource5 int32  `json:"cost_res_5" db:"cost_res_5"`
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
	VillageID     int32  `json:"village_id" db:"village_id"`
	PlayerID      int32  `json:"player_id" db:"player_id"`
	VillageName   string `json:"village_name" db:"village_name"`
	VillageSize   int32  `json:"village_size" db:"village_size"`
	VillageStatus int32  `json:"village_status" db:"village_status"`
	VillageLocY   int32  `json:"village_loc_y" db:"village_loc_y"`
	VillageLocX   int32  `json:"village_loc_x" db:"village_loc_x"`
}

// HunterHut     int64 `json:"hunter_hut" db:"hunter_hut"`
// 	WoodcutterHut int64 `json:"woodcutter_hut" db:"woodcutter_hut"`
// 	Quarry        int64 `json:"quarry" db:"quarry"`
// 	CopperMine    int64 `json:"copper_mine" db:"copper_mine"`
// 	Fountain      int64 `json:"fountain" db:"fountain"`

type VillageSetup struct {
	VillageID       int32 `json:"village_id" db:"village_id"`
	PlayerID        int32 `json:"player_id" db:"player_id"`
	HunterHut_1     int32 `json:"hunterhut_1" db:"hunterhut_1"`
	HunterHut_2     int32 `json:"hunterhut_2" db:"hunterhut_2"`
	HunterHut_3     int32 `json:"hunterhut_3" db:"hunterhut_3"`
	HunterHut_4     int32 `json:"hunterhut_4" db:"hunterhut_4"`
	HunterHut_5     int32 `json:"hunterhut_5" db:"hunterhut_5"`
	WoodcutterHut_1 int32 `json:"woodcutterhut_1" db:"woodcutterhut_1"`
	WoodcutterHut_2 int32 `json:"woodcutterhut_2" db:"woodcutterhut_2"`
	WoodcutterHut_3 int32 `json:"woodcutterhut_3" db:"woodcutterhut_3"`
	WoodcutterHut_4 int32 `json:"woodcutterhut_4" db:"woodcutterhut_4"`
	WoodcutterHut_5 int32 `json:"woodcutterhut_5" db:"woodcutterhut_5"`
	Quarry_1        int32 `json:"quarry_1" db:"quarry_1"`
	Quarry_2        int32 `json:"quarry_2" db:"quarry_2"`
	Quarry_3        int32 `json:"quarry_3" db:"quarry_3"`
	Quarry_4        int32 `json:"quarry_4" db:"quarry_4"`
	Quarry_5        int32 `json:"quarry_5" db:"quarry_5"`
	CopperMine_1    int32 `json:"coppermine_1" db:"coppermine_1"`
	CopperMine_2    int32 `json:"coppermine_2" db:"coppermine_2"`
	CopperMine_3    int32 `json:"coppermine_3" db:"coppermine_3"`
	CopperMine_4    int32 `json:"coppermine_4" db:"coppermine_4"`
	CopperMine_5    int32 `json:"coppermine_5" db:"coppermine_5"`
	Fountain_1      int32 `json:"fountain_1" db:"fountain_1"`
	Fountain_2      int32 `json:"fountain_2" db:"fountain_2"`
	Fountain_3      int32 `json:"fountain_3" db:"fountain_3"`
	Fountain_4      int32 `json:"fountain_4" db:"fountain_4"`
	Fountain_5      int32 `json:"fountain_5" db:"fountain_5"`
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
