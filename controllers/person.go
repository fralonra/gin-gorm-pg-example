package controllers

import (
	"../db"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func GetPersonAll(c *gin.Context) {
	var people []models.Person
	db.GetDb().Find(&people)
	c.JSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	var person models.Person
	id := c.Param("id")
	db.GetDb().Where("id = ?", id).First(&person)
	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  db.GetDb().Create(&person)
  c.JSON(http.StatusOK, &person)
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")

	db := db.GetDb()
  if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }
  c.BindJSON(&person)
  db.Save(&person)
  c.JSON(http.StatusOK, &person)
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")
	if personID, err := strconv.Atoi(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
    return
  } else {
		person.ID = personID
	}

	db.GetDb().Delete(&person)
	c.JSON(http.StatusOK, &person)
}