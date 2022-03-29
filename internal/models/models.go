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

type Production struct {
	HunterHut     int64 `json:"hunter_hut" db:"hunter_hut"`
	WoodcutterHut int64 `json:"woodcutter_hut" db:"woodcutter_hut"`
	Quarry        int64 `json:"quarry" db:"quarry"`
	CopperMine    int64 `json:"copper_mine" db:"copper_mine"`
	Fountain      int64 `json:"fountain" db:"fountain"`
}

type Resource struct {
	Food   int64 `json:"food" db:"food"`
	Wood   int64 `json:"wood" db:"wood"`
	Stone  int64 `json:"stone" db:"stone"`
	Copper int64 `json:"copper" db:"copper"`
	Water  int64 `json:"water" db:"water"`
}

// resource_id INTEGER PRIMARY KEY,
// resource TEXT NOT NULL,
// quality INTEGER NOT NULL,
// rate INTEGER NOT NULL,

type ResourceModel struct {
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
