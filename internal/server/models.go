package server

type Player struct {
	PlayerID    string `json:"player_id"`
	PlayerName  string `json:"player_name"`
	PlayerScore int64  `json:"player_score"`
	Active      bool   `json:"active"`
}
