package main

import (
	"SCRAPING-INFORMATIKAUMM/config"
	"SCRAPING-INFORMATIKAUMM/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	schedule := router.Group("/news")
	{
		schedule.GET("/", inDB.GetAllNews)
	}

	err := router.Run(":4591")

	if err != nil {
		panic("Error when running router")
	}
}