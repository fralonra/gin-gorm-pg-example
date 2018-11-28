package controllers

import (
	"../db"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func GetPlayerAll(c *gin.Context) {
	var players []models.Player
	db.GetDb().Find(&players)
	c.JSON(http.StatusOK, players)
}

func GetPlayerByID(c *gin.Context) {
	var player models.Player
	id := c.Param("id")
	db.GetDb().Where("id = ?", id).First(&player)
	c.JSON(http.StatusOK, player)
}

func CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.BindJSON(&player); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  db.GetDb().Create(&player)
  c.JSON(http.StatusOK, &player)
}

func UpdatePlayer(c *gin.Context) {
	var player models.Player
	id := c.Param("id")

	db := db.GetDb()
  if err := db.Where("id = ?", id).First(&player).Error; err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }
  c.BindJSON(&player)
  db.Save(&player)
  c.JSON(http.StatusOK, &player)
}

func DeletePlayer(c *gin.Context) {
	var player models.Player
	id := c.Param("id")
	if playerID, err := strconv.Atoi(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
    return
  } else {
		player.ID = playerID
	}

	db.GetDb().Delete(&player)
	c.JSON(http.StatusOK, &player)
}