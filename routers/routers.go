package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) {
	router.GET("/", controllers.Home)
	router.GET("/players", controllers.PlayersView)
	router.GET("/player/:id", controllers.PlayerView)
	
	api := router.Group("/api")
	{
		player := api.Group("/player")
		player.GET("/", controllers.GetPlayerAll)
		player.GET("/:id", controllers.GetPlayerByID)
		player.POST("/", controllers.CreatePlayer)
		player.PUT("/:id", controllers.UpdatePlayer)
		player.DELETE("/:id", controllers.DeletePlayer)
	}
}
