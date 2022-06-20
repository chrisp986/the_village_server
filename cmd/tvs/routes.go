package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *application) routes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		// v1.GET("/players", getPlayers)
		v1.GET("/player/:player_id", a.getPlayer)
		// v1.GET("/active_players", getActivePlayers)

		//1. create new player and then create a new village that belongs to the player
		v1.POST("/new_player", a.postPlayer)
		v1.POST("/train_worker", a.postTrainNewWorker)

	}
	web := r.Group("/")
	{
		web.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main TV website",
			})
		})
	}
}
