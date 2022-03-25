package player

import "github.com/chrisp986/the_village_server/internal/server"

func () insertNewPlayer(newPlayer server.Player) error {

	ConnectDB()

	_, err = db.NamedExec(`
		INSERT INTO players (player_id, player_name, player_score, active, connected)
		VALUES (:player_id, :player_name, :player_score, :active, :connected)
		ON CONFLICT(player_name) DO UPDATE SET
		player_id = :player_id,
		player_name = :player_name,
		player_score = :player_score,
		active = :active,
		connected = :connected
		`, newPlayer)
	if err != nil {
		return err
	}
	return nil

}
