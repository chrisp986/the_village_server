package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/players", getPlayers)

		v1.GET("/", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.String(http.StatusOK, "Welcome Gin Server")
		})
	}
}

var players = []Player{
	{PlayerID: "1", PlayerName: "John Doe", PlayerScore: 100, Active: true},
	{PlayerID: "2", PlayerName: "Jane Doe", PlayerScore: 200, Active: true},
}

func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}
