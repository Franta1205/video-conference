package main

import (
	"github.com/gin-gonic/gin"
	"video-conference/routes"
)

func main() {
	router := gin.Default()

	routes.SetUpRoutes(router)
	router.Run(":8080")
}
