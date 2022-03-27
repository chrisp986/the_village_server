package main

import (
	"net/http"
	"strconv"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/gin-gonic/gin"
)

var players = []models.Player{
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

func getPlayer(c *gin.Context) {

	playerID := c.Param("player_id")
	pID, err := strconv.Atoi(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var player models.Player
	for _, p := range players {
		if p.PlayerID == int32(pID) {
			player = p
		}
	}

	if player.PlayerID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No player found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, player)
}

func postPlayers(c *gin.Context) {
	var newPlayer models.Player

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

	var numProductions models.Production

	if err := c.BindJSON(&numProductions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newResources := models.Resource{
		Food:   numProductions.HunterHut * 2,
		Wood:   numProductions.WoodcutterHut * 2,
		Stone:  numProductions.Quarry * 2,
		Copper: numProductions.CopperMine * 2,
		Water:  numProductions.Fountain * 2,
	}

	c.IndentedJSON(http.StatusOK, newResources)
}
