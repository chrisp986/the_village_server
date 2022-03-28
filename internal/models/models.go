package models

import "time"

type Player struct {
	PlayerID       int32     `json:"player_id"`
	PlayerName     string    `json:"player_name"`
	PlayerEmail    string    `json:"player_email"`
	PlayerPassword string    `json:"player_password"`
	PlayerScore    int64     `json:"player_score"`
	Active         bool      `json:"active"`
	Connected      bool      `json:"connected"`
	Created        time.Time `json:"created"`
}

type Production struct {
	HunterHut     int64 `json:"hunter_hut"`
	WoodcutterHut int64 `json:"woodcutter_hut"`
	Quarry        int64 `json:"quarry"`
	CopperMine    int64 `json:"copper_mine"`
	Fountain      int64 `json:"fountain"`
}

type Resource struct {
	Food   int64 `json:"food"`
	Wood   int64 `json:"wood"`
	Stone  int64 `json:"stone"`
	Copper int64 `json:"copper"`
	Water  int64 `json:"water"`
}
