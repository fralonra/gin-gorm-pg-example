package main

import (
	"./config"
	"./db"
	"./routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	initConfig()
	initDb()
	initServer()
}

func initConfig() {
	if err := config.LoadConfig("config/config.yml"); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func initDb() {
	db.Init()
}

func initServer() {
	serverConfig := config.GetServer()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.LoadHTMLGlob("views/*")
	routers.RegisterRouters(router)
	router.Run(serverConfig.Port)
}
