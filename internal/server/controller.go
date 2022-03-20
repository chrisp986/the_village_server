package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var players = []Player{
	{
		PlayerID:    1,
		PlayerName:  "player1",
		PlayerScore: 0,
		Active:      true,
		Connected:   false,
	},
	{
		PlayerID:    2,
		PlayerName:  "player2",
		PlayerScore: 0,
		Active:      true,
		Connected:   true,
	},
}

func getPlayers(c *gin.Context) {
	if len(players) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No players found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, players)
}

func getPlayer(c *gin.Context) {}

func getActivePlayers(c *gin.Context) {}

func postPlayers(c *gin.Context) {
	var newPlayer Player

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	players = append(players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

func postCalculateNewResources(c *gin.Context) {

	var numProductions Production

	if err := c.BindJSON(&numProductions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newResources := Resource{
		Food:   numProductions.HunterHut * 2,
		Wood:   numProductions.WoodcutterHut * 2,
		Stone:  numProductions.Quarry * 2,
		Copper: numProductions.CopperMine * 2,
		Water:  numProductions.Fountain * 2,
	}

	c.IndentedJSON(http.StatusOK, newResources)
}
