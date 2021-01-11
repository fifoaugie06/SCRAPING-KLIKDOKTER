package main

import (
	"SCRAPING-KLIKDOKTER/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	infoSehat := router.Group("/info-sehat")
	{
		infoSehat.GET("/", controllers.GetInfoSehat)
	}

	err := router.Run(":4591")

	if err != nil {
		panic("Error when running router")
	}
}
