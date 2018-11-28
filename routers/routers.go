package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) {
	router.GET("/", controllers.Home)
	router.GET("/people", controllers.PeopleView)
	router.GET("/person/:id", controllers.PersonView)
	
	api := router.Group("/api")
	{
		person := api.Group("/person")
		person.GET("/", controllers.GetPersonAll)
		person.GET("/:id", controllers.GetPersonByID)
		person.POST("/", controllers.CreatePerson)
		person.PUT("/:id", controllers.UpdatePerson)
		person.DELETE("/:id", controllers.DeletePerson)
	}
}
