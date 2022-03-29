package main

import (
	"net/http"
	"strconv"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/gin-gonic/gin"
)

// func getPlayers(c *gin.Context) {
// 	if len(players) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No players found!"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, players)
// }

func (a *application) getPlayer(c *gin.Context) {

	playerID := c.Param("player_id")
	pID, err := strconv.Atoi(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := a.players.Get(pID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (a *application) postPlayers(c *gin.Context) {
	var newPlayer models.Player

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	id, err := a.players.Insert(newPlayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"player_id": id,
	})
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
