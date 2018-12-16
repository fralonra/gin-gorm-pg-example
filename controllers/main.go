package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/people")
}

func PeopleView(c *gin.Context) {
	c.HTML(http.StatusOK, "people.html", gin.H{
		"title": "People",
	})
}

func PersonView(c *gin.Context) {
	c.HTML(http.StatusOK, "person.html", gin.H{
		"title": "Person",
	})
}
