package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"video-conference/routes"
)

func main() {
	router := gin.Default()

	routes.SetUpRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
