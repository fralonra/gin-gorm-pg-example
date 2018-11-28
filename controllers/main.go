package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/players")
}

func PlayersView(c *gin.Context) {
	c.HTML(http.StatusOK, "players.html", gin.H{
		"title": "Players",
	})
}

func PlayerView(c *gin.Context) {
	c.HTML(http.StatusOK, "player.html", gin.H{
		"title": "Player",
	})
}